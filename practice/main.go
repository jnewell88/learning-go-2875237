package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	//dow := rand.Intn(7) + 1
	//fmt.Println("Day", dow)
	//fallthrough can be used to continue through the case,otherwise it will exit after one is met.
	var result string
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "It's Sunday!"
		//fallthrough
	case 2:
		result = "It's Monday!"
		//fallthrough
	default:
		result = "It's some other day!"
	}
	fmt.Println(result)
}
