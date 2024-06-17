// Progam to show the use of nested if statement
package main
import "fmt"
func main() {
	//taking two local-variable
	var d1 int = 700
	var d2 int = 390
	// using if statement
	if(d1 == 700 ){
		// if condition is true then
		// check the following
		if (d2 == 800 ){
			// if the condition is true
			// then displayh following

			fmt.Printf("Value of d1 is 700 and d2 is 390\n");
		}
	}
}