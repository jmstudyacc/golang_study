package main

import "fmt"

func main() {
	// A slice expression is a slice of a slice - this is similar in how Python operates
	x := []int{1, 2, 3, 4}
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
	fmt.Println("Changing an element in a slice affects all slices that share that element:")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// It becomes even worse when you start using append()
	x = []int{1, 2, 3, 4}
	y = x[:2]
	fmt.Println()
	fmt.Println("Appending with slices can create a lot of confusion when the result is output:")
	fmt.Println("Capacity of 'x':", cap(x),"\nCapacity of 'y':",cap(y))
	y = append(y, 30)
	fmt.Println("\nx:", x)
	fmt.Println("y:", y)
	fmt.Println()

	// The chaos can run out of control very quickly
	x = make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y = x[:2]
	z = x[2:]
	fmt.Println("Capacity of 'x':",cap(x), "\nCapacity of 'y':",cap(y), "\nCapacity of 'z':", cap(z))
	fmt.Println()
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("The result of too much unsafe slice manipulation:")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	// Needless to say this is fairly chaotic, so it is best to use a full slice expression if you need this functionality
	x = make([]int, 0, 5)
	x = append(x, 1, 2, 3, 4)
	y = x[:2:2]
	z = x[2:4:4]
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println()
	fmt.Println("Full slice expression is used, lines 60 & 61, which results in a more predictable outcome:")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

}

