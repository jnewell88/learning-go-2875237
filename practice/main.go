package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum (before rounding):", floatSum)

	//Managing precision
	//if initializing, use :=, if changing value, use =

	//Round to nearest int
	//floatSum = math.Round(floatSum)

	//Round to 1 decimal (round to nearest)
	//No matter if you use 1,10,100,1000... its always 1 decimal place
	//floatSum = math.Round(floatSum*100) / 100

	//** BEST OPTION **
	//This is another option - (round down)
	//This actually follows with # of decimal places (10 - 1, 100 - 2, 1000 - 3....)
	floatSum = math.Floor(floatSum*1000) / 1000

	//And another - (round up)
	//No matter if you use 1,10,100,1000... its always 1 decimal place
	//floatSum = math.Ceil(floatSum*100) / 100
	fmt.Println("The sum is now", floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi

	fmt.Printf("Given a circle with radius %.2f, the circumference is: %.2f\n", circleRadius, circumference)

}
