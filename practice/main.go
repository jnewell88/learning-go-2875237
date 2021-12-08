package main

import (
	"fmt"
)

func main() {
	//Classes/structs, data and methods
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 9
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
}

//Custom type
//lowercase = non exported (private), uppercase = exported (public)
type Dog struct {
	Breed  string
	Weight int
}
