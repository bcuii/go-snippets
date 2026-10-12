//go:build ignore

// Multiple senders: who closes the channel?
// No sender can close it: the others may still send (panic: send on closed channel),
// or another sender closes it again (panic: close of closed channel).
// Instead, wait until all senders are done, then close exactly once.
//
// sync.WaitGroup counts the running senders: Add(1) before starting one,
// Done() when it finishes, Wait() blocks until the count is back to 0.
//
// Run: go run ./channel/patterns/02_multi_sender_close.go
package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := range 3 { // 3 senders
		wg.Add(1)
		go func() {
			ch <- i
			wg.Done()
		}()
	}

	// In its own goroutine: if main waited here, the senders would block on ch,
	// because nobody is receiving yet.
	go func() {
		wg.Wait() // all senders are done
		close(ch) // so closing once is safe
	}()

	for v := range ch { // order varies: the senders run at the same time
		fmt.Println("got", v)
	}
}
