//go:build ignore

// Arguments of a deferred call are evaluated at the defer statement.
// A deferred closure reads variables when it runs.
//
// Run: go run ./defer/basic/03_args_vs_closure.go
package main

import "fmt"

func main() {
	x := 1
	defer fmt.Println("arg:", x)                  // x evaluated now: 1
	defer func() { fmt.Println("closure:", x) }() // x read at return: 2
	x = 2
}
