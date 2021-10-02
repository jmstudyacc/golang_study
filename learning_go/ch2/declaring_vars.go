package main

// Avoid assigning variables here as it cna be difficult to track the changes made to it

import "fmt"

func main() {
	// Declaration via keyword, explicit type and assignment
	var a int = 10
	fmt.Printf("Variable 'a' is of type: %T\n", a)

	// Go will infer the variable type based on the assignment, so you can drop the type declaration
	var b = 20
	fmt.Printf("Variable 'b' is of type: %T\n", b)

	// Alternatively you may want to declare an empty variable - this WILL need a declaration (there's nothing to infer)
	var c int
	fmt.Printf("Variable 'c' is of type: %T\n", c)
	// If you have multiple variables of the same type you can assign them together - positional assignments
	var d, e = 30, 40
	fmt.Printf("Variable 'd' is of type: %T, Variable 'e' is of type: %T\n", d, e)
	// This concept is consistent with zero values as well
	var f, g int
	fmt.Printf("Variable 'a' is of type: %T, Variable 'g' is of type: %T\n", f, g)

	// You can even assign different variable types together
	var h, i = 50, "Xin chào"
	fmt.Printf("Variable 'h' is of type: %T, Variable 'i' is of type: %T\n", h, i)

	// You can include declaration lists to express multiple variables at once
	var (
		j    int
		k        = 60
		l    int = 70
		m, n     = 80, "Cám ơn"
		o, p string
	)

	fmt.Printf("\nVariable 'j' is of type: %T\n", j)
	fmt.Printf("Variable 'k' is of type: %T\n", k)
	fmt.Printf("Variable 'l' is of type: %T\n", l)
	fmt.Printf("Variable 'm' is of type: %T, Variable 'n' of type: %T\n", m, n)
	fmt.Printf("Variable 'o' is of type: %T, Variable 'p' is of type: %T\n", o, p)

	// Go also has a short declaration syntax
	q := 90
	fmt.Printf("\nVariable 'q' is of type: %T", q)

	// You can also use this syntax to assign to a pre-existing variable, provided at least one new variable is on the left
	q, r := 100, "Tạm biệt"
	fmt.Printf("Variable 'q' is of type: %T, Variable 'r' is of type: %T\n", q, r)
}
