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

	float1, float2, operation := GetInputs()
	result := 0.0
	switch operation {
	case "+":
		fmt.Println("Adding....")
		result = Add(float1, float2)
	case "-":
		fmt.Println("Subtracting....")
		result = Subtract(float1, float2)
	case "/":
		fmt.Println("Dividing....")
		result = Divide(float1, float2)
	case "*":
		fmt.Println("Multiplying....")
		result = Multiply(float1, float2)
	default:
		panic("Invalid Operation")
	}
	fmt.Println("The result is", math.Round(result*100)/100)
}

func GetInputs() (value1 float64, value2 float64, operation string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value 1: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	fmt.Print("Value 2: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}
	fmt.Print("Select an operation (+ - * /): ")
	input3, _ := reader.ReadString('\n')
	operator := strings.TrimSpace(input3)
	return float1, float2, operator
}

//Could alternatively make single GetInput function and just call it twice

func Add(value1 float64, value2 float64) float64 {
	return value1 + value2
}
func Subtract(value1, value2 float64) float64 {
	return value1 - value2
}
func Divide(value1, value2 float64) float64 {
	if value2 == 0 {
		panic("Can't divide by zero!")
	}
	return value1 / value2
}
func Multiply(value1, value2 float64) float64 {
	return value1 * value2
}
