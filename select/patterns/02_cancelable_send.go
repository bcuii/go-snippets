//go:build ignore

// Cancelable send: a producer selects on both out <- v and ctx.Done().
// When the consumer stops reading and cancels, the producer returns instead of blocking on the send forever.
// Any goroutine that sends to a channel nobody may read anymore needs such a way out.
//
// Run: go run ./select/patterns/02_cancelable_send.go
package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for v := range generate(ctx) {
		if v == 3 {
			break // the consumer needs only 0, 1, 2
		}
		fmt.Println("got", v)
	}
	cancel()
	time.Sleep(10 * time.Millisecond)                  // give the generator a moment to return
	fmt.Println("goroutines:", runtime.NumGoroutine()) // 1: only main is left
}

// generate sends 0, 1, 2, ... until ctx is canceled.
func generate(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; ; i++ {
			select {
			case out <- i:
			case <-ctx.Done(): // without this case, the goroutine blocks on out <- 4 forever
				return
			}
		}
	}()
	return out
}
