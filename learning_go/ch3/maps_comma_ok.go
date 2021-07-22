package main

import "fmt"

func main() {
	// If you request a value for a key that is not in the map the map returns 0
	dogAges := map[string]int{}
	dogAges["Paddy"] = 4
	dogAges["Rupert"] = 2
	dogAges["Winnie"] = 1
	fmt.Println(dogAges["Pancake"])

	// However sometimes you will want to see if a key is in the map, not if it has a 0 value
	//dogAges["Pancake"] = 0

	v, ok := dogAges["Paddy"]
	fmt.Println(v, ok)
	fmt.Println()

	// 'ok' returns a bool where false = not in the map - this is used often for 0 value vs. not in map
	v, ok = dogAges["Pancake"]
	fmt.Println(v, ok)
}
