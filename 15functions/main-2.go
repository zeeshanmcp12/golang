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

	result := strings.TrimSpace(userInput())
	fmt.Println("Your input:", result)

	fmt.Printf("Type of input: %T", result)

	// fmt.Println()
	// fmt.Printf("Tsype of input %T", userInput())

}

func userInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter any text: ")
	input, _ := reader.ReadString('\n')
	return input
}

func strToInt(strText string) int {
	input, err := strconv.Atoi(strText)
	if err != nil {
		panic(err)
	}
	return input

}
