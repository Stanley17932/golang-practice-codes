package main

import "fmt"

func main() {
	const C = "Pop"
	var D = "PiiksofPiiks"

	// Concat the strings.
	var helloEveryone = C + " " + D

	// Print the result."
	fmt.Println(helloEveryone)

	// Compare the Strings
	fmt.Println(C == "POP")
	fmt.Println(D < C)

}
