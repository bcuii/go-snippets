//go:build ignore

// Embedding: Dog embeds Animal, so Animal's fields and methods are promoted to Dog.
// It is composition, not inheritance:
//   - a Dog is not an Animal: it can't be used where an Animal is expected
//   - no overriding: Animal.Hello calls Animal.Name, even when called through a Dog
//
// For polymorphism, use an interface.
//
// Run: go run ./struct/basic/05_embedding.go
package main

import "fmt"

type Animal struct{ Legs int }

func (a Animal) Name() string  { return "animal" }
func (a Animal) Hello() string { return "I am " + a.Name() }

type Dog struct {
	Animal // embedded: the field name is the type name, d.Animal
}

func (d Dog) Name() string { return "dog" } // hides Animal.Name for d.Name()

func main() {
	d := Dog{Animal{Legs: 4}}
	fmt.Println(d.Legs)    // 4: same as d.Animal.Legs
	fmt.Println(d.Name())  // dog
	fmt.Println(d.Hello()) // I am animal: d.Hello() is d.Animal.Hello(), its receiver is the Animal
	// var a Animal = d // compile error: cannot use d (variable of struct type Dog) as Animal value
}
