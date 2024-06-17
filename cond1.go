// A progtam to demonstrate the use of the if statement

package main
import "fmt"

func main(){
	// taking the local variable

	var d int =600

	// using if statement for checking the condition
	if (d <1000) {
		// print following if
		// condition evaluates to true
		fmt.Printf("d is less than 1000\n")
	}
	fmt.Printf("Value of d is: %d\n", d)
}