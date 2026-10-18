//go:build ignore

// defer runs when the function returns, not at the end of a loop iteration.
// In a loop, deferred calls pile up and everything stays open until the function returns.
// Fix: move the loop body into its own function.
//
// Run: go run ./defer/basic/05_loop.go
package main

import "fmt"

func main() {
	leaky()
	fmt.Println("---")
	fixed()
}

func leaky() {
	for i := 1; i <= 3; i++ {
		fmt.Println("open", i)
		defer fmt.Println("close", i) // all 3 stay open until leaky returns
	}
	fmt.Println("loop done")
}

func fixed() {
	for i := 1; i <= 3; i++ {
		handle(i)
	}
	fmt.Println("loop done")
}

func handle(i int) {
	fmt.Println("open", i)
	defer fmt.Println("close", i) // runs when handle returns: once per iteration
}
