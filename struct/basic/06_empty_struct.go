//go:build ignore

// struct{} has no fields and takes 0 bytes. Use it when only presence matters, not a value:
// a set, a signal on a channel, a context key type (see context/basic/04_with_value.go).
//
// Run: go run ./struct/basic/06_empty_struct.go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(struct{}{})) // 0: struct{} is the type, struct{}{} is its value

	set := map[string]struct{}{} // set: only the keys matter
	set["a"] = struct{}{}
	_, ok := set["a"]
	fmt.Println("has a:", ok)

	// Signal: done carries no data, it only says "the work is finished".
	// With chan bool, readers would ask what true and false mean; chan struct{} says there is no value.
	// ctx.Done() returns <-chan struct{} for the same reason.
	done := make(chan struct{})
	go func() {
		fmt.Println("working")
		close(done) // close wakes every receiver; no value needs to be sent
	}()
	<-done // blocks until done is closed
	fmt.Println("done")
}
