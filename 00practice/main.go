package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Reading notes to rephrase the learning!")
	fmt.Println("Taking input from user!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter any text: ")
	input, _ := reader.ReadString('\n')
	converted, _ := strconv.Atoi(strings.TrimSpace(input))
	fmt.Printf("Type of string is %T and value is %v", converted, input)

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
