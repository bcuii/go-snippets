//go:build ignore

// Struct literal and zero value. Fields left out of a literal get their zero value,
// so a struct is usable without a constructor (e.g. var mu sync.Mutex, var buf bytes.Buffer).
// If every field is comparable, structs can be compared with == and used as map keys.
//
// Run: go run ./struct/basic/01_literal_zero.go
package main

import "fmt"

type Point struct {
	X, Y int
}

type Tagged struct {
	Name string
	Tags []string // a slice can't be compared, so Tagged can't either
}

func main() {
	var zero Point      // all fields zero
	p := Point{X: 1}    // named fields: Y is 0
	q := Point{1, 2}    // positional: breaks when a field is added; prefer names
	ptr := &Point{X: 3} // pointer to a new Point
	fmt.Println("zero:", zero)
	fmt.Println("p:   ", p)
	fmt.Println("q:   ", q)
	fmt.Println("*ptr: ", *ptr)
	fmt.Println("ptr: ", ptr)
	fmt.Printf("%+v\n", p) // %+v also prints field names

	fmt.Println(p == Point{X: 1}) // true: compares field by field
	seen := map[Point]bool{p: true}
	fmt.Println(seen[Point{X: 1}]) // true

	// Tagged{} == Tagged{}   // compile error: struct containing []string cannot be compared
	// map[Tagged]bool{}      // compile error: invalid map key type Tagged
}
