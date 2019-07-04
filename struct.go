package main

import "fmt"

// Person - structs are typed collections of fields.
// They’re useful for grouping data together to form records.
// This person struct type has name and age fields.
type Person struct {
	name   string
	weight float32
	height float32
	age    int
}

func showPerson(p Person) {
	fmt.Println("Nome: ", p.name)
	fmt.Println("Idade: ", p.age)
	fmt.Println("Peso: ", p.weight)
	fmt.Println("Altura: ", p.height)
}

func main() {
	// p := Person{name: "Gabriel", age: 18}
	// This syntax creates a new struct. and you can name the fields
	// when initializing a struct
	x := Person{name: "asas", age: 19, height: 170.0, weight: 72.0}
	showPerson(x)

	// Structs are mutable.
	s := Person{name: "Sean", age: 50}
	// fmt.Println(s.name)
	sp := &s
	sp.age = 52
	fmt.Println(s.age)
}
