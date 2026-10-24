//go:build ignore

// recover stops a panic, but only when called directly by a deferred function.
// After recovering, the function returns normally to its caller. It does not resume at the panic.
//
// Run: go run ./panic/basic/02_recover.go
package main

import "fmt"

func main() {
	safe()
	fmt.Println("main: still running")

	notDirect()
	// never reached: the panic escaped
}

func safe() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()
	panic("boom")
	// never reached: safe returns right after the defer
}

func notDirect() {
	defer func() {
		logPanic() // recover is one call deeper here, so it returns nil
	}()
	panic("boom again")
}

func logPanic() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}
