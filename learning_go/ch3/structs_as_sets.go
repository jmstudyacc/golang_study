package main

import "fmt"

func main() {
	// Structs can provide a memory improvement over maps as an empty struct uses 0 bytes
	intSet := map[int]struct{}{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10 }

	// Using a Struct for a set makes the code more clumsy than a map
	for _ , v := range vals {
		intSet[v] = struct{}{}
	}

	// Comma ok idiom is needed to check if a value is in the set
	if _, ok := intSet[5]; ok {
		fmt.Println("5 is in the set")
	}
}
