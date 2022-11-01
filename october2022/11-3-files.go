package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files!")

	// os.Create
	// io.WriteString
	// ioutil.ReadFile

	content := "Hello world"

	file, err := os.Create("./11-3-file.txt")

	checkNilErr(err)

	io.WriteString(file, content)

	defer file.Close()
	readFile("./11-3-file.txt")

}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	fmt.Println("Here is the data from file:", string(data))

	checkNilErr(err)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
