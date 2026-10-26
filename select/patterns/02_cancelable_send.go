//go:build ignore

// Cancelable send: a sender whose reader may leave selects on both out <- v and ctx.Done().
// When the reader stops and cancels, the sender returns instead of blocking on the send forever.
//
// Run: go run ./select/patterns/02_cancelable_send.go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for p := range pages(ctx) {
		fmt.Println("got page", p)
		if p == 2 {
			break // found what we need: stop reading
		}
	}
	cancel()                           // tell pages to stop
	time.Sleep(100 * time.Millisecond) // wait to see pages stop
}

// pages fetches pages from an API in the background; there are 100, the reader needs 2.
func pages(ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer fmt.Println("pages: stopped")
		defer close(out) // lets a reader that reads all 100 pages end its range
		for p := 1; p <= 100; p++ {
			select {
			case out <- p:
			case <-ctx.Done(): // without this case, pages blocks on out <- 3 forever
				return
			}
		}
	}()
	return out
}
