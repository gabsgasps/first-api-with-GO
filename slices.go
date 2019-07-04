package main

import (
	"fmt"
)

func main() {
	// To create an empty slice with non-zero length,
	// use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued)
	s := make([]string, 3)

	// fmt.Println(s)

	// We can set and get just like with arrays.
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	// fmt.Println("set:", s)
	// fmt.Println("get:", s[2])

	// len returns the length of the slice as expected.
	// fmt.Println("len:", len(s))

	// c := make([]string, len(s))
	// copy(c, s)
	// fmt.Println(c)

	// In addition to these basic operations, slices support several more that make them richer than arrays.
	// One is the builtin append, which returns a slice containing one or more new values.
	s = append(s, "d", "e", "f", "g", "h", "i")
	// fmt.Println("set:", s)

	// slice[low:high]
	// l := s[2:5]
	// fmt.Println(l)

	// This slices up to (but excluding) s[2].
	// l := s[:2]
	// fmt.Println("sl2:", l)

	// And this slices up from (and including) s[2].
	// l := s[2:]
	// fmt.Println("sl2:", l)

	// tw := []string{"2", "4", "5", "9"}
	l := s[3:]
	fmt.Println(l)
}
