//go:build ignore

// WithCancelCause (Go 1.20): cancel with a reason.
// ctx.Err() stays context.Canceled, so errors.Is checks keep working; context.Cause(ctx) returns the reason.
//
// Run: go run ./context/basic/05_cause.go
package main

import (
	"context"
	"errors"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancelCause(context.Background())
	cancel(errors.New("upstream closed"))

	fmt.Println("err:  ", ctx.Err())          // context canceled
	fmt.Println("cause:", context.Cause(ctx)) // upstream closed
}
