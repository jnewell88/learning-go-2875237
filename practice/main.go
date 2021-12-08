package main

import (
	"fmt"
	"sort"
)

func main() {
	//Slice is a non initialized size array
	//Array
	//var colors = [3]string{"Red", "Green", "Blue"}
	//Slice
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	//Append color to the end
	colors = append(colors, "Purple")
	//remove the first color in the slice
	colors = append(colors[1:len(colors)])
	fmt.Println(colors)
	//remove the last color in the slice
	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	//Specified size
	//numbers := make([]int, 5, 5)
	//Dynamic size
	numbers := make([]int, 5)
	numbers[0] = 134
	numbers[1] = 54
	numbers[2] = 65
	numbers[3] = 88
	numbers[4] = 555
	fmt.Println(numbers)

	//Dynamic size required to append
	numbers = append(numbers, 235)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
