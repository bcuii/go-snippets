//go:build ignore

// When several cases are ready, select picks one at random, not the first one written.
//
// Run: go run ./select/basic/02_random.go
package main

import "fmt"

func main() {
	a := make(chan int)
	b := make(chan int)
	close(a) // a closed channel is always ready to receive
	close(b)

	countA, countB := 0, 0
	for range 1000 {
		select {
		case <-a:
			countA++
		case <-b:
			countB++
		}
	}
	fmt.Println("a:", countA, "b:", countB) // roughly half each
}
