package main

import "fmt"

func main() {
	// Go doesn't have sets, but it allows for maps to simulate Sets
	// You use the key of the map for the type and a bool for the value
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	// Iterate over values in vals with a for-range loop
	for _, v := range vals {
		// Place the element into the intSet map with a bool value of true
		intSet[v] = true
	}

	// Length of the map will be shorter, you cannot have duplicate keys in a map
	fmt.Println("Length of vals:", len(vals), "\nLength of intSet:", len(intSet))
	fmt.Println(intSet)
	fmt.Println()

	// Prints 'true' as 5 is in the map
	fmt.Println(intSet[5])

	// Prints 'false' as 500 is NOT in the map
	fmt.Println(intSet[500])


	if intSet[100] {
		fmt.Println("100 is in the set")
	} else {
		fmt.Println(intSet[100])
	}
}
