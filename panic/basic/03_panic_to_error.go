//go:build ignore

// Turn a panic into an error: recover in a defer and assign the named result.
// Common at boundaries (handlers, callbacks, plugins) so one bad call doesn't crash everything.
//
// Run: go run ./panic/basic/03_panic_to_error.go
package main

import "fmt"

func main() {
	bad := func() {
		var p *int
		*p = 1 // runtime panic: nil pointer dereference
	}
	fmt.Println("named:  ", safeCall(bad))
	fmt.Println("unnamed:", safeCallUnnamed(bad)) // <nil>: the panic is lost
}

// err must be a named result: after a panic no return statement runs,
// so only the defer can set the return value.
func safeCall(f func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r) // becomes the return value
		}
	}()
	f()
	return nil
}

// Wrong: unnamed result. The panic is recovered, but the caller gets nil.
func safeCallUnnamed(f func()) error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r) // sets the local only
		}
	}()
	f()
	return err // never runs after the panic
}
