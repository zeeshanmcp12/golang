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
	content := userInput()

	file, err := os.Create("./userData.txt")

	io.WriteString(file, content)

	CheckNilErr(err)
	readFile("./userData.txt")

	defer file.Close()

}

func userInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter any text to add in file: ")
	input, _ := reader.ReadString('\n')
	f_text := strings.TrimSpace(input)
	return f_text

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileData string) {
	data, _ := ioutil.ReadFile(fileData)
	fmt.Print("File to be written with data\n", string(data))

}
