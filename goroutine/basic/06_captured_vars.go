//go:build ignore

// A variable captured by a goroutine outlives the function that created it, whatever its type.
//
// Run: go run ./goroutine/basic/06_captured_vars.go
package main

import (
	"fmt"
	"time"
)

func main() {
	start()
	fmt.Println("start returned")
	time.Sleep(100 * time.Millisecond) // let the goroutine print
}

func start() {
	n := 42
	names := []string{"a", "b"}
	go func() {
		time.Sleep(50 * time.Millisecond) // start has returned by now
		fmt.Println(n, names)             // still alive: captured by the goroutine
	}()
}
