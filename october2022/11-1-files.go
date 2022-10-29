package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Working with files in golang!")

	content := "This content came from main function."

	file, err := os.Create("./11-file.txt")
	checkNilErr(err)

	io.WriteString(file, content)

	defer file.Close()

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
