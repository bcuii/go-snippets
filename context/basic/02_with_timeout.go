//go:build ignore

// WithTimeout: the ctx cancels itself after the duration; Err() is DeadlineExceeded.
// WithDeadline is the same but takes a time.Time: WithTimeout(d) = WithDeadline(time.Now().Add(d)).
// Unlike time.After, the ctx can be passed down, so every call below shares one deadline.
// Still defer cancel(): it frees the timer if the work finishes early.
//
// Run: go run ./context/basic/02_with_timeout.go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(time.Second): // a slow operation
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("gave up:", ctx.Err())
	}
}
