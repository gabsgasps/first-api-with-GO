package main

import "fmt"

func f(from string) {
	fmt.Println(from)
}

func main() {
	f("izi")
	go f("opa")

	// fmt.Scanln()
	fmt.Println("done")
}
