//go:build ignore

// panic stops the current function and runs its defers, then does the same in each caller,
// up to the top of the goroutine. If nobody recovers, the program crashes (exit status 2).
//
// Run: go run ./panic/basic/01_unwind.go
package main

import "fmt"

func main() {
	defer fmt.Println("main: defer runs") // 3rd
	a()
	// never reached
}

func a() {
	defer fmt.Println("a: defer runs") // 2nd
	b()
	// never reached
}

func b() {
	defer fmt.Println("b: defer runs") // 1st
	panic("something went wrong")      // unwinding starts here
}
