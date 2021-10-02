package main

import "fmt"

func main() {
	// there are a number of built-in functions provided by Go

	// we'll start by looking at len()
	var x = []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	fmt.Println("Current length is:", len(x)) // len can work on slices and arrays of any types
	fmt.Println("Current capacity is:", cap(x))
	fmt.Println()

	// append is next and you'll use this to grow slices only
	x = append(x, 6)
	fmt.Println(x)
	fmt.Println("Current length is:", len(x)) // as you can see 6 is added to the end of the slice
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
	fmt.Println()

	// It is not possible to create an empty with length and capacity defined without using MAKE
	x = make([]int, 5)
	fmt.Println("This creates a slice of length 5 and capacity 5, with all elements initialized to 0.")
	fmt.Println(x)

	// Specifying the capacity is as easy as including the parameter in the MAKE function call
	x = make([]int, 5, 10)
	fmt.Println("This creates a slice of length 5 and capacity 10, with the elements initialized to 0")
	fmt.Println(x)

	// Go community is divided on how to declare a slice, but I will be going with the following
	x = make([]int, 0, 10)
	fmt.Println("This creates a slice of length 0 and capacity 10, with no elements present")
	fmt.Println(x, "Length:", len(x), "Capacity:", cap(x))
	fmt.Println("You need to use append() to add elements to this slice")
	x = append(x, 1, 2, 3, 4, 5)
	fmt.Println(x, "Length:", len(x), "Capacity:", cap(x))

}
