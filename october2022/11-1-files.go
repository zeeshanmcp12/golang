package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Working with files in golang!")

	content := "This content came from main function."

	file, err := os.Create("./11-file.txt")
	checkNilErr(err)

	io.WriteString(file, content)

	defer file.Close()

	readFile("./11-file.txt")

}

func readFile(filename string) {
	fmt.Println("Please find below the text: ")
	data, err := ioutil.ReadFile(filename)
	checkNilErr(err)
	fmt.Println(string(data))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
