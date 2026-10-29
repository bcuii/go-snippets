//go:build ignore

// Structs are values: assignment and function arguments copy every field.
// The copy is shallow: a slice, map or pointer field copies only the pointer to its data,
// so both copies share that data.
//
// Run: go run ./struct/basic/02_copy.go
package main

import "fmt"

type Team struct {
	Name    string
	Members []string
}

func main() {
	a := Team{Name: "x", Members: []string{"bob"}}
	b := a               // copies every field
	b.Name = "y"         // a.Name is still "x"
	b.Members[0] = "amy" // a sees this too: both Members point to the same array
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	rename(a)
	fmt.Println(a.Name) // still "x": rename got a copy; pass *Team to change a
}

func rename(t Team) { t.Name = "z" }
