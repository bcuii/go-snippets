//go:build ignore

// Multiple defers run in LIFO order: last deferred, first run.
// So resources are released in reverse order of acquisition.
//
// Run: go run ./defer/basic/02_lifo.go
package main

import "fmt"

func main() {
	fmt.Println("lock A")
	defer fmt.Println("unlock A") // runs second

	fmt.Println("lock B")
	defer fmt.Println("unlock B") // runs first
}
