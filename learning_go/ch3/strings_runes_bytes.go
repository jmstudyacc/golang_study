package main

import "fmt"

func main() {
	// Indexing into a String is possible, index starts at 0,
	var s string = "Hello there"
	var b byte = s[6]
	// This will return a byte value, not a letter
	fmt.Println(b)
	// This converts it to a String to print out the letter - only do this if you know it is only 1 byte long
	fmt.Println(string(b))

	s2 := s[4:7]
	s3 := s[:5]
	s4 := s[6:]

	fmt.Println("s:", s)
	fmt.Println("s2:", s2, "\ns3:", s3, "\ns4:", s4)

	// Do not slice or index, instead extract substrings and code points using functions in
	// strings and unicode/utf8 packages
}
