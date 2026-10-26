//go:build ignore

// Goroutine leak: the caller stops waiting, but the sender is blocked on ch <- v forever.
// An unbuffered send needs a receiver, and the only one (fetch's select) has left.
// Fix for a one-shot result: make(chan T, 1). The send goes into the buffer, so the goroutine finishes.
//
// Run: go run ./select/patterns/01_goroutine_leak.go
package main

import (
	"fmt"
	"time"
)

func main() {
	fetch(0)                           // unbuffered
	fetch(1)                           // buffer 1
	time.Sleep(300 * time.Millisecond) // wait to see which sender finishes
}

// fetch waits up to 100ms for a slow call that takes 200ms.
func fetch(size int) {
	ch := make(chan string, size) // created here, but captured by the goroutine: outlives fetch
	go func() {
		time.Sleep(200 * time.Millisecond) // slow call
		ch <- "result"                     // unbuffered: stuck here forever
		fmt.Println("cap", cap(ch), "sender done")
	}()
	select {
	case <-ch:
	case <-time.After(100 * time.Millisecond):
		fmt.Println("cap", cap(ch), "timeout")
	}
}
