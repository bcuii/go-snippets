//go:build ignore

// Limit concurrency: a buffered channel works as a semaphore.
// Its capacity is the max number of tasks running at once. Send to take a slot, receive to give it back.
// Use it when one goroutine per item could overload something (CPU, a DB, a remote API).
// Outside the standard library, errgroup.SetLimit (golang.org/x/sync) does the same.
//
// Run: go run ./goroutine/patterns/01_limit_concurrency.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	sem := make(chan struct{}, 3) // at most 3 at a time
	var wg sync.WaitGroup

	for i := 1; i <= 7; i++ {
		sem <- struct{}{} // take a slot: blocks the loop while 3 are running
		wg.Go(func() {
			defer func() { <-sem }() // give the slot back
			fmt.Println(time.Since(start).Round(100*time.Millisecond), "start task", i)
			time.Sleep(100 * time.Millisecond) // do the work
		})
	}
	wg.Wait()
}
