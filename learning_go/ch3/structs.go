package main

import "fmt"

func main() {
	/*
	Maps are convenient but they do not define an API as there is no way to constrain a map to contain certain keys
	Furthermore, all the values in the map must be of the same type - maps are not ideal for passing data from function to function
	Related data that should be grouped means you should define a struct
	*/

	// This struct, in its current format, is defined in main() so it is only usable within main()
	type person struct {
		name string
		age int
	}

	// Once a Struct is declared you can define variables of that type
	var fred person

	// No values have been assigned to 'fred' so 0 is printed
	fmt.Println(fred)

	// You can also assign a struct literal to a variable
	bob := person{}

	// Again, no values, 0 is printed
	fmt.Println(bob)

	// Assigning values to the struct fields will be visible when printing the person variable
	james := person {
		name: "James",
		age: 29,
	}
	fmt.Println(james)

	// You access the values of a struct via dotted notation
	bob.name = "Bob"
	fmt.Println(bob.name)

}

