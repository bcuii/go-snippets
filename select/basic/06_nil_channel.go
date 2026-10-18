//go:build ignore

// A nil channel is never ready, so setting a channel to nil disables its case.
// Common use: merge channels and turn each one off once it is closed.
// Without nil, a closed channel is always ready and the loop would spin on zero values.
//
// Run: go run ./select/basic/06_nil_channel.go
package main

import "fmt"

func main() {
	a := make(chan int, 2)
	a <- 1
	a <- 2
	close(a)

	b := make(chan int, 1)
	b <- 10
	close(b)

	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				a = nil // closed: disable this case
				continue
			}
			fmt.Println("a:", v)
		case v, ok := <-b:
			if !ok {
				b = nil
				continue
			}
			fmt.Println("b:", v)
		}
	}
	fmt.Println("both closed")
}
