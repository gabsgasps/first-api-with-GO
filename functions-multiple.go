package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func main() {
	// a, b := vals()

	// If you only want a subset of the returned values, use the blank identifier _.
	// c, _ := vals()
	_, c := vals()
	fmt.Println(c)
}
