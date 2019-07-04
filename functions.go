package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	// res := plus(1, 2)

	res := plusPlus(1, 2, 3)
	fmt.Println(res)

}
