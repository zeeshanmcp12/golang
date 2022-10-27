package main

import (
	"fmt"
	"io"
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
	content := "Working with files in golang!"

	// Create file
	file, err := os.Create("./10-files-txt-file.txt")
	checkNilErr(err)

	// Write string into file
	io.WriteString(file, content)

	// Make sure to close the file after writing/reading
	defer file.Close()

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
