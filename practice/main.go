package main

import (
	"fmt"
)

func main() {
	anInt := 42
	//& = memory address
	//* = memory value
	var p = &anInt
	fmt.Println("Value of p:", *p)

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Value 1:", *pointer1)

	//Setting pointer to a new value, which changes value1
	*pointer1 = *pointer1 / 31
	fmt.Println("Pointer 1:", *pointer1)
	fmt.Println("Value 1:", value1)
}
