package main

import (
	"fmt"
)

func main() {
	doSomething()
	sum := addValues(5, 8)
	fmt.Println("The sum is:", sum)
	multiSum, multiCount := addAllValues(4, 7, 9, 11)
	fmt.Println("Sum of multiple values:", multiSum)
	fmt.Println("Count of items", multiCount)
}

//Not returning anthing, don't need a return type beween () and {

func doSomething() {
	fmt.Println("Doing something")
}

func addValues(value1 int, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
