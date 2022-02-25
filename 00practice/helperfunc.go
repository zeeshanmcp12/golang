package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Helper Function!")
	fmt.Printf("Enter any number: ")
	result := userInput()
	fmt.Println()
	fmt.Printf("Type of %v after conversion is: %T ", result, result)
}

func userInput() int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Printf("Type of input before conversion is: %T ", text)
	return strtoint(text)
}

func strtoint(value string) int {
	convertToInt, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		fmt.Println(err)
	}
	return convertToInt
}
