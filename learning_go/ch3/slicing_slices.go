package main

import "fmt"

func main() {
	// A slice expression is a slice of a slice - this is similar in how Python operates
	x := []int{1, 2, 3, 4, 5}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)
	fmt.Println()

	// A slice of a slice is not a copy of data - it shares the memory of the original variable & unused capacity
	x[1] = 20
	y[0] = 10
	z[1] = 30

	// changing 'x' modifies 'y' and 'z' AND changing 'y' and 'z' modified 'x' - this can easily become confusing
	fmt.Println("Changing an element in a slice affects all slices that share that element.")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// Needless to say this is fairly chaotic, so it is best to use a full slice expression if you need this functionality

}

