//go:build ignore

// WithCancel: derive a ctx and a cancel func. Calling cancel closes ctx.Done().
// Cancellation is cooperative: the worker must watch ctx.Done() and return by itself.
// Always call cancel (usually defer cancel()), or the ctx leaks; go vet's lostcancel check catches this.
//
// Run: go run ./context/basic/01_with_cancel.go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Go(func() {
		for {
			select {
			case <-ctx.Done(): // closed by cancel()
				fmt.Println("worker stopped:", ctx.Err())
				return
			case <-time.After(100 * time.Millisecond):
				fmt.Println("working")
			}
		}
	})

	time.Sleep(350 * time.Millisecond)
	cancel()
	wg.Wait()
}
