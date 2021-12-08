package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("A simple calculator")

	fmt.Print("Value 1: ")
	numInput1, _ := reader.ReadString('\n')
	aFloat1, err := strconv.ParseFloat(strings.TrimSpace(numInput1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	fmt.Print("Value 2: ")

	numInput2, _ := reader.ReadString('\n')
	aFloat2, err := strconv.ParseFloat(strings.TrimSpace(numInput2), 64)

	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}

	valSum := aFloat1 + aFloat2
	valSum = math.Round(valSum*100) / 100
	//%v is the placeholder for using the default formatting based on type of variable
	fmt.Printf("The sum of %v and %v is %v\n", aFloat1, aFloat2, valSum)

}
