package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Working with files in golang!")

	// Taking user input by calling userInput() funcion
	content := userInput()

	// Creating file
	file, err := os.Create("./userData.txt")

	// Writing content to file
	io.WriteString(file, content)

	CheckNilErr(err)

	// Reading data from the file created above
	readFile("./userData.txt")

	// Closing file at the end of execution
	defer file.Close()

}

// function to take input from user
func userInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter any text to add in file: ")
	input, _ := reader.ReadString('\n')
	f_text := strings.TrimSpace(input)
	return f_text

}

// Check if err has something in it
func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

// ReadFile function
func readFile(fileData string) {
	data, _ := ioutil.ReadFile(fileData)
	fmt.Print("File to be written with data\n", string(data))

}
