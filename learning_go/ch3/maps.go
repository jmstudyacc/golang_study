package main

import "fmt"

func main() {
	fmt.Println("Testing out maps in Go:")
	// map type = map[keyType]valueType | maps can be declared similarly to variables
	var nilMap map[string]int
	fmt.Println("Nil map:", nilMap)

	// Here an empty map literal is being assigned - different to a nil map, reading it will return 0
	totalWins := map[string]int{}
	fmt.Println("Empty map literal:", totalWins)
	fmt.Println()

	// Example of a nonempty map literal
	teams := map[string][]string {
		"Orcas": []string{"Fred", "Ralph", "Bijou"},
		"Lions": []string{"Sarah", "Peter", "Billie"},
		"Kittens": []string{"Waldo", "Raul", "Ze"},
	}
	fmt.Println(teams)

	// Passing a map to len() tells the number of key-value pairs in the list
	fmt.Println("Number of key-value pairs in 'teams' variable:",len(teams))
	fmt.Println()

	// If you know the number of k, v pairs you will use but not the values, you can use make to make a map
	ages := make(map[int][]string, 10)
	fmt.Println("'ages' map:",ages)
	fmt.Println()

	// Slices = lists that are processed sequentially
	// Maps = organized data not strictly in increasing order (use a map when order is not important)

	// Reading & Writing a Map!
	fmt.Println("The following examples show how you can read from a map:")
	teamWins := map[string]int{}
	teamWins["Orcas"] = 1
	teamWins["Lions"] = 2
	fmt.Println("Orcas wins:", teamWins["Orcas"])
	fmt.Println("Kittens wins:", teamWins["Kittens"])
	teamWins["Kittens"]++
	fmt.Println()
	fmt.Println("Kittens wins:", teamWins["Kittens"])
	teamWins["Lions"] = 3
	fmt.Println("Lions wins:", teamWins["Lions"])
	fmt.Println()

	// Deleting from Maps is pretty simple and is done by the delete() function
	m := map[string]int {
		"hello": 5,
		"world": 10,
	}
	fmt.Println(m)
	delete(m, "hello")
	fmt.Println(m)



}
