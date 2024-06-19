package main

import "fmt"

type dog struct {
	name  string
	breed string
}

func main() {
	pet := dog{
		name:  "Maximus",
		breed: "Rottweiler",
		age:   5,
	}

	fmt.Printf("%+v", pet)
}
