package main

import "fmt"

func main() {

	// for is Go’s only looping construct. Here are three basic types of for loops.
	// i := 1
	// The most basic type, with a single condition.
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }
	// A classic initial/condition/after for loop.
	for j := 0; j <= 5; j++ {
		fmt.Println("yeah", j)
	}

	// You can also continue to the next iteration of the loop.
	// for n := 0; n <= 5; n++ {
	//     if n%2 == 0 {
	//         continue
	//     }
	//     fmt.Println(n)
	// }

}
