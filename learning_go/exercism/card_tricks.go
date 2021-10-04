package main

import "fmt"

func VerifyIndex(slice []int, index int) bool {
	// if the length of the slice - 1 (indexing) is greater than the index and index isn't negative
	if len(slice)-1 >= index && index > -1 {
		// we verify that the index value passed is valid
		return true
	} else {
		return false
	}
}

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index existed in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	// checking the index value is valid
	if VerifyIndex(slice, index) {
		// it is, so return the value in the slice at the indexed value and true
		return slice[index], true
	}
	// return 0 and false to indicate an error occurred
	return 0, false
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range it is be appended.
func SetItem(slice []int, index, value int) []int {
	if VerifyIndex(slice, index) {
		// returning the slice with the new value passed by user in the index position thanks to helper function
		slice[index] = value
		return slice
	}
	// otherwise return the slice with the required value appended to the end
	return append(slice, value)
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	// clumsily create an empty slice
	res := []int{}

	// loop a number of time = to length passed in
	for i := 0; i < length; i++ {
		// on each iteration append the user value to the slice
		res = append(res, value)
	}
	// finish up and return the 'filled' slice
	return res
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if VerifyIndex(slice, index) {
		// golang has a weird way of removing things from slices?
		slice = append(slice[:index], slice[index+1:]...)
	}
	return slice
}

func main() {
	card, ok := GetItem([]int{1, 2, 4, 1}, 2)
	fmt.Println(card)
	// Output: 4
	fmt.Println(ok)
	// Output: true

	fmt.Println()

	deck, status := GetItem([]int{1, 2, 4, 1}, 10)
	fmt.Println(deck)
	// Output: 0
	fmt.Println(status)
	// Output: false

	idx := -1
	val := 6
	cards := []int{1, 2, 4, 1}
	res := SetItem(cards, idx, val)
	// Output: [1 2 4 1 6]
	fmt.Println(res)

	cards = []int{5, 2, 10, 6, 8, 7, 0, 9}
	val = 1
	idx = 4
	res2 := SetItem(cards, idx, val)
	// Output: [1 2 6 1]
	fmt.Println(res2)

	fmt.Println(PrefilledSlice(10, 2))

	fmt.Println(RemoveItem([]int{3, 2, 6, 4, 8}, 2))
	// Output: []int{3, 2, 4, 8}
}
