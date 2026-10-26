//go:build ignore

// WithValue: attach request-scoped data (request ID, trace ID, auth user) to a ctx.
// Value looks up the chain from child to parent; a missing key returns nil.
// Don't use it for optional parameters: pass those as normal arguments.
//
// Run: go run ./context/basic/04_with_value.go
package main

import (
	"context"
	"fmt"
)

// Usually each package defines its own unexported key type:
// keys compare by type and value, so no other package can collide.
type ctxKey string

const requestIDKey ctxKey = "requestID"

func main() {
	ctx := context.WithValue(context.Background(), requestIDKey, "req-42")
	ctx, cancel := context.WithCancel(ctx) // derived ctx still sees the value
	defer cancel()

	handle(ctx)
	handle(context.Background())
	fmt.Println(ctx.Value("requestID")) // <nil>: a plain string is a different key
}

func handle(ctx context.Context) {
	id, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		id = "unknown"
	}
	fmt.Println("request id:", id)
}
