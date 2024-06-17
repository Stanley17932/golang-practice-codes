package main
import "fmt"

func main(){
	c :=6
	// Using address of operator(&) and
	// pointer indirection(*) operator
	d := &c 
	fmt.Println(*d)

	*d = 7
	fmt.Println(c)
}