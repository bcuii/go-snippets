//go:build ignore

// The most common use of directional channels: a function creates a channel,
// sends and closes it in its own goroutine, and returns only the receive end.
// Callers can receive, but can't send or close. time.After and ctx.Done() work this way.
//
// It pays off when values are not available right away (I/O, events, time).
//
// Run: go run ./channel/patterns/01_direction_generator.go
package main

import "fmt"

func generate(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range n {
			out <- i
		}
		close(out)
	}()
	return out
}

func main() {
	nums := generate(3)
	for v := range nums {
		fmt.Println("got", v)
	}
	// nums <- 1   // compile error: cannot send to receive-only channel
	// close(nums) // compile error: cannot close receive-only channel
}
