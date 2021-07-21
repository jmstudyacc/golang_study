package main

import "fmt"

// working with slices is like working with arrays but without needing to specify a size

func main() {
	var x = []int{10, 20, 30} // using [...] makes an array [] makes a slice
	fmt.Printf("\nThis prints out a slice with 3 values:\n%d\n", x)

	var y = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Printf("\nThis prints out a slice that places '0' in undeclared positions:\n%d\n", y)

	var z [][]int
	fmt.Printf("\nThis prints out an empty slice with no values:\n%d\n", z)

	fmt.Printf("\nThis prints out the value held in var y at index 5:\n%d\n", y[5])

	// as you can see this is very similar to using arrays, but differences abound!
	// for instance Slices are not comparable and will produce a compile time error
	fmt.Printf("\nThis prints out the result of the statement, 'x == nil':\n%t\n", x == nil) // slices can ONLY be compared against nil
}
