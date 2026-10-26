//go:build ignore

// Batch: flush when the batch is full, or when time is up, whichever comes first.
// Common for logs, metrics and bulk inserts. Real code also flushes the remaining items when the input closes.
//
// Run: go run ./select/patterns/04_batch.go
package main

import (
	"fmt"
	"time"
)

func main() {
	in := make(chan int)
	go func() {
		for i := 1; i <= 4; i++ {
			in <- i
		}
	}()

	var batch []int
	timeUp := time.After(200 * time.Millisecond)
	for {
		select {
		case v := <-in:
			batch = append(batch, v)
			if len(batch) == 3 {
				fmt.Println("full:   ", batch)
				batch = nil
			}
		case <-timeUp:
			fmt.Println("time up:", batch) // don't wait forever for a full batch
			return
		}
	}
}
