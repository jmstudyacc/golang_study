package main

import "fmt"

func main() {
	// arrays are rarely used in Go but do exist - Declaration type 1 begins by specifying the size of the array - the values will equal 0
	var x [3]int
	fmt.Println(x)

	// declaration type 2 begins by specifying the size and the values in the array
	var y = [3]int{10, 20, 30}
	fmt.Println(y)

	// declaration 3 is a bit weirder
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}

	// this creates an array of 12 ints with the following
	// values: [1, 0, 0, 0, 0, 4, 6, 0, 0, 0, 100, 15]
	fmt.Println(z)

	// you can also use an array literal to init an array and leave off the number and use ... instead
	var xyz = [...]int{30, 20, 10}
	fmt.Println(xyz)

	// arrays can be compared with boolean operators
	fmt.Println(x == y) // returns false
	fmt.Println(x != y) // returns true

	// arrays in Go are 1-D but you can simulate multidimensional arrays
	// notice no variable assignment
	var a [2][3]int
	fmt.Println(a) // this is not a true matrix - it is an array of 2 with value of arrays of length 3 ints

	// arrays are as usual read with square brackets
	fmt.Println(xyz[1])

	// the built-in function len takes in an ARRAY and returns its length
	fmt.Println(len(x))

	// Go deviates from the norm and makes an array become an array of type size
	// so [3]int array is a different type to an array that is declared as [4]int
	// Go array sizes cannot be specified by a variable as types must be resolved at compile time
}
