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
	fmt.Println("Working with Files")
	content := userInput()

	file, err := os.Create("./newUserData.txt")
	CheckNilErr(err)

	io.WriteString(file, content)

	readFile()
	defer file.Close()

}

func readFile() {
	data, err := ioutil.ReadFile("./newUserData.txt")
	CheckNilErr(err)
	fmt.Print("Data to be written in file\n", string(data))
}

func userInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter any text to write in file: ")
	input, _ := reader.ReadString('\n')
	f_text := strings.TrimSpace(input)
	return f_text

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
