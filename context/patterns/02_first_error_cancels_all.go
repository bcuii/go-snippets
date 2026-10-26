//go:build ignore

// First error cancels all: run workers in parallel; the first one that fails stops the rest.
// This is the core of errgroup.WithContext (golang.org/x/sync).
//
// Run: go run ./context/patterns/02_first_error_cancels_all.go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	ctx, cancel := context.WithCancelCause(context.Background())
	defer cancel(nil)

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Go(func() {
			if err := work(ctx, i); err != nil {
				cancel(err) // only the first call sets the cause; a stopped worker's ctx.Err() changes nothing
			}
		})
	}
	wg.Wait()
	fmt.Println("first error:", context.Cause(ctx))
	fmt.Println("elapsed:", time.Since(start).Round(10*time.Millisecond)) // not 1s
}

func work(ctx context.Context, id int) error {
	if id == 2 {
		time.Sleep(50 * time.Millisecond)
		return fmt.Errorf("worker %d failed", id)
	}
	select {
	case <-time.After(time.Second): // normal work takes 1s
		fmt.Println("worker", id, "done")
		return nil
	case <-ctx.Done():
		fmt.Println("worker", id, "stopped:", context.Cause(ctx))
		return ctx.Err()
	}
}
