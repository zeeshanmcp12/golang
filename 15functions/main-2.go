package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Handling user input using function")

	fmt.Printf("Enter any text: ")

	result := strings.TrimSpace(userInput())
	fmt.Printf("Input: %v, Type: %T", result, result)

	fmt.Println()

	fmt.Print("Enter any number: ")
	text := userInput()

	fmt.Printf("Type of text is %T\n", text)

	fmt.Printf("Input: %v, Type: %T", strToInt(text), strToInt(text))

}

func userInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func strToInt(strText string) int {
	input, err := strconv.Atoi(strings.TrimSpace(strText))
	// if err != nil {
	// 	panic(err)
	// }
	checkErr(err)
	return input

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
