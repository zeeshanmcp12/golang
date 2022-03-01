package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files in golang")
	content := "The text written in file created by golang."
	file, err := os.Create("./newfile.txt")
	CheckNilErr(err)

	textLength, err := io.WriteString(file, content)
	CheckNilErr(err)
	fmt.Println("Text length is: ", textLength)
	readFile()

	// data, _ := ioutil.ReadFile("./gofile.txt")

	defer file.Close()
}

func readFile() {
	data, _ := ioutil.ReadFile("./newfile.txt")
	fmt.Println("The data we read from file is:\n", string(data))

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
