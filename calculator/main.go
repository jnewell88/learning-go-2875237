package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from Go!"
	file, err := os.Create("./fromString.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote a file with %v characters\n", length)
	//Defer: wait until everything else has executed, then run this
	defer file.Close()
	//File needs to be completely closed before trying to read
	defer readFile("./fromString.txt")
}

func readFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	checkError(err)
	fmt.Println("Text read from file:", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
