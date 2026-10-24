//go:build ignore

// Deadlock from lock order: g1 locks a then b, g2 locks b then a.
// Each holds one lock and waits for the other forever.
// Fix: every goroutine takes the locks in the same order (always a before b).
//
// The runtime only reports it because main is blocked too. In a server where other
// goroutines keep running, the two would just hang silently.
//
// Run: go run ./sync/basic/03_lock_order_deadlock.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var a, b sync.Mutex
	var wg sync.WaitGroup

	wg.Go(func() {
		a.Lock()
		fmt.Println("g1: locked a, waiting for b")
		time.Sleep(100 * time.Millisecond) // make the bad timing certain; in real code it happens by chance
		b.Lock()
		// never reached
	})
	wg.Go(func() {
		b.Lock()
		fmt.Println("g2: locked b, waiting for a")
		time.Sleep(100 * time.Millisecond)
		a.Lock()
		// never reached
	})

	wg.Wait() // every goroutine is blocked: fatal error: all goroutines are asleep - deadlock!
}
