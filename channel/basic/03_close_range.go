//go:build ignore

// close + range: after close, range drains the remaining values and then exits.
// Without close, range drains the values and then waits forever for more.
//
// range knows when to stop because receiving from a closed, drained channel
// returns the zero value immediately, with ok == false.
//
// Run: go run ./channel/basic/03_close_range.go
package main

import "fmt"

func main() {
	// With close
	buf := make(chan int, 2)
	buf <- 11
	buf <- 33

	close(buf)
	for v := range buf {
		fmt.Println("range got", v)
	}
	fmt.Println("range exited")

	// Without close
	noClose := make(chan int, 2)
	noClose <- 11
	noClose <- 33

	for v := range noClose {
		fmt.Println("range (no close) got", v)
	}
	// never reached: range keeps waiting for more values
}
