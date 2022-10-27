package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Create, Write and Read the file in golang!")

	// Following packages/utilities are going to be used here
	// os.Create
	// io.WriteString
	// file.Close()
	// ioutil.ReadFile

	// First, content
	content := "This content needs to be added/updated in txt file!"

	// Create file
	file, err := os.Create("./10-files-txt-file.txt")
	checkNilErr(err)

	// Write string into file
	length, err := io.WriteString(file, content)
	checkNilErr(err)

	fmt.Println("Length of content is", length)

	// Make sure to close the file after writing/reading
	defer file.Close()

	// Function call to read the file
	readFile("./10-files-txt-file.txt")

}

// Reading file
func readFile(filename string) {
	data, _ := ioutil.ReadFile(filename)
	fmt.Println("Content is \n", string(data))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
