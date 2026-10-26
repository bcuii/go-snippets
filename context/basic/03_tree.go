//go:build ignore

// Cancel tree: every derived ctx is a child of its parent.
// Canceling a parent cancels all its descendants; canceling a child leaves the parent alone.
//
// Run: go run ./context/basic/03_tree.go
package main

import (
	"context"
	"fmt"
)

func main() {
	parent, cancelParent := context.WithCancel(context.Background())
	a, cancelA := context.WithCancel(parent)
	b, cancelB := context.WithCancel(parent)
	defer cancelB() // always call cancel, even if the parent will cancel it first

	cancelA()
	fmt.Println("after cancel a:")
	fmt.Println("  parent:", parent.Err()) // <nil>
	fmt.Println("  a:     ", a.Err())      // context canceled
	fmt.Println("  b:     ", b.Err())      // <nil>

	cancelParent()
	fmt.Println("after cancel parent:")
	fmt.Println("  parent:", parent.Err())
	fmt.Println("  b:     ", b.Err()) // canceled through its parent
}
