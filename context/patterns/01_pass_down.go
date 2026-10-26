//go:build ignore

// Pass ctx down: every function that may block takes ctx as its first parameter and passes it on.
// The top sets the deadline once; every layer below stops when it expires.
//
// Scenario: an API must answer within 150ms and runs several queries; they all share that budget.
// A 100ms timeout per query could take 300ms, and results after the caller gives up are wasted.
//
// Real APIs work the same way: db.QueryContext(ctx, ...), http.NewRequestWithContext(ctx, ...).
//
// Run: go run ./context/patterns/01_pass_down.go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
	defer cancel()

	fmt.Println("handle:", handle(ctx))
}

func handle(ctx context.Context) error {
	for _, table := range []string{"users", "orders", "items"} {
		if err := query(ctx, table); err != nil {
			return fmt.Errorf("query %s: %w", table, err)
		}
	}
	return nil
}

// query stands in for a DB or HTTP call that accepts a ctx.
func query(ctx context.Context, table string) error {
	select {
	case <-time.After(100 * time.Millisecond): // each query takes 100ms
		fmt.Println("got", table)
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
