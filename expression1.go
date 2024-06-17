// Program to show
// Concept of the variable

package main

import "fmt"

func main() {
	// using short variable declarations
	// short variable declaration acts as
	// an assignment for the my_var2 variable
	// because the variable is present in the same block

	my_var2, my_var3 := 39, 56
	my_var1, my_var4 := 56, 100

	fmt.Printf("Value of my_var1 and my_var4 is: %d%d\n", my_var1, my_var4)
	fmt.Printf("Value of my_var2 and my_var3 is: %d%d\n", my_var2, my_var3)

}
