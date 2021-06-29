package main

import "fmt"

// working with slices is like working with arrays but without needing to specify a size

func main() {
	var x = []int{10, 20, 30} // using [...] makes an array [] makes a slice
	fmt.Println(x)

	var y = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(y)

	var z [][]int
	fmt.Println(z)

	fmt.Println(y[5])

	// as you can see this is very similar to using arrays, but differences abound!
	// for instance Slices are not comparable and will produce a compile time error
	fmt.Println(x == nil) // slices can ONLY be compared against nil
}
