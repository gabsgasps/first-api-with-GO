package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	// for num := 0; num < len(nums); num++ {
	// 	fmt.Println(nums[num])
	// }
	for _, num := range nums {
		sum += num
	}

	// range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "opa", "b": "beleza"}

	for k, v := range kvs {
		fmt.Println(k, v)
	}

	// fmt.Println(sum)
}
