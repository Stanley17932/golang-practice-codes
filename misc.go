package main
import "fmt"
func main(){
	var c int = 46
	var d int = 85

	// "=" (Simple Assignment)
	c=d
	fmt.Println(c)
	// "+=" (Add assignment)
	c +=d
	fmt.Println(c)

	// "-=" (Subtract Assignment)

	c -=d 
	fmt.Println(c)

	// "*=" (Multiply Assingment)
	c*=d 
	fmt.Println(c)
	// "/=" (Division Assignment)
	c /=d
	fmt.Println(c)
	// "%=" (Modulus Assignment)
	c %=d
	fmt.Println(c)
}