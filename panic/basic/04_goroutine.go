//go:build ignore

// recover only catches panics in its own goroutine.
// A panic in another goroutine crashes the whole program, even if main recovers.
// So every goroutine that may panic needs its own defer + recover.
//
// Run: go run ./panic/basic/04_goroutine.go
package main

import (
	"fmt"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("main recovered:", r) // never runs: the panic is in another goroutine
		}
	}()

	go func() {
		panic("panic in goroutine")
	}()

	time.Sleep(time.Second) // wait for the goroutine
	// never reached: the program crashes during the sleep
}
