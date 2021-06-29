package main

import "fmt"

func main() {
	// there are a number of built-in functions provided by Go
	// we'll start by looking at len()
	var x = []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println("Current length is:", len(x))	// len can work on slices and arrays of any types
	fmt.Println("Current capacity is:", cap(x))
	fmt.Println()

	// append is next and you'll use this to grow slices only
	x = append(x, 6)
	fmt.Println(x)
	fmt.Println("Current length is:", len(x))	// as you can see 6 is added to the end of the slice
	fmt.Println("Current capacity is:", cap(x))
	fmt.Println()

	// append can also append multiple values to a slice
	x = append(x, 7, 8, 9, 10)
	fmt.Println(x)
	fmt.Println("Current length is:", len(x))
	fmt.Println("Current capacity is:", cap(x))
	fmt.Println()

	// you can also append a slice onto another slice
	var y = []int{11, 12, 13, 14, 15}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("Current length is:", len(x))
	fmt.Println("Current capacity is:", cap(x))

	// auto-sizing is inefficient and this can be addressed by the use of make

	fmt.Println()
	z := make([]int, 5)
	fmt.Println(z)

}

