package main

import "fmt"

type Dog struct{
	Breed string
}

func main(){
	kind :=Dog{
		Breed: "Rottweiler",
	}
	fmt.Println(kind.Breed)
}