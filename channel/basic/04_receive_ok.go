//go:build ignore

// v, ok := <-ch in four states. ok == false only when the channel is closed and drained.
//
//	   state             | v          | ok    | blocks?
//	---------------------+------------+-------+-----------------
//	1. open, has data    | the value  | true  | no
//	2. closed, has data  | the value  | true  | no
//	3. closed, drained   | zero value | false | no
//	4. open, empty       | -          | -     | yes, until a send
//
// Run: go run ./channel/basic/04_receive_ok.go
package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// 1. Open, has data: get the value, ok == true.
	ch <- 11
	v, ok := <-ch
	fmt.Println("open, has data:  ", v, ok)

	// 2. Closed, still has data: get the value, ok == true.
	ch <- 33
	close(ch)
	v, ok = <-ch
	fmt.Println("closed, has data:", v, ok)

	// 3. Closed, drained: zero value immediately, ok == false.
	v, ok = <-ch
	fmt.Println("closed, drained: ", v, ok)

	// 4. Open, empty: blocks until someone sends. Nobody does, so it waits forever.
	empty := make(chan int)
	fmt.Println("open, empty:      waiting...")
	v, ok = <-empty
	// never reached
}
