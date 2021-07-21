package main

import "fmt"

func main() {
	// Go does not allow for automatic type promotion between variables - this requires type conversion
	var x int32 = 10
	fmt.Printf("Type of variable 'x': %T\n", x)

	var y float64 = 30.2
	fmt.Printf("Type of variable 'y': %T\n", y)

	var z int64 = 20
	fmt.Printf("Type of variable 'z': %T\n", z)

	// The line below this will not compile due to incompatible types
	//var xy = x + y

	// This line will compile as explicit type conversion is used - note that decimal accuracy is lost in this conversion
	var xy = x + int32(y)
	fmt.Printf("\nThe value of xy: %d\nThe type of xy: %T", xy, xy)

}