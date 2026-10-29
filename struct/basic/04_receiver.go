//go:build ignore

// Value receiver: the method gets a copy, so changes are lost.
// Pointer receiver: the method gets the address, so changes stay.
// If any method needs a pointer receiver, use pointer receivers for all methods of the type.
//
// Use a pointer receiver when:
//   - the method changes the receiver
//   - the struct has a field that must not be copied, like sync.Mutex (go vet: copylocks)
//   - the struct is large, so copying it on every call costs too much
//
// Use a value receiver for:
//   - small types that don't change, like time.Time, Point, type Celsius float64
//   - map, chan and func types: they already refer to their data
//
// If unsure, use a pointer receiver.
// Source: https://go.dev/wiki/CodeReviewComments#receiver-type
//
// Run: go run ./struct/basic/04_receiver.go
package main

import "fmt"

type Counter struct{ n int }

func (c Counter) IncV()  { c.n++ } // changes a copy
func (c *Counter) IncP() { c.n++ } // changes the original

func main() {
	var c Counter
	c.IncV()
	fmt.Println(c.n) // 0
	c.IncP()         // Go takes the address for you: same as (&c).IncP()
	fmt.Println(c.n) // 1
	// Counter{}.IncP() // compile error: cannot call pointer method IncP on Counter (a literal has no address)
}
