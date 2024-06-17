package main

import "fmt"

const Pie = 3.14

func main() {
	var radius float64
	fmt.Println("Enter the radius of the circle")
	fmt.Scan(&radius)
	area := Pie * radius * radius
	fmt.Println("Area of the circle is", area)
}
