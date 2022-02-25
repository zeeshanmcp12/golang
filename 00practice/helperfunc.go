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
	fmt.Printf("Enter any text: ")
	result := userInput()
	fmt.Println(result)
	fmt.Printf("Type of result is %T", result)
}

func userInput() int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strtoint(text)
}

func strtoint(value string) int {
	convertToInt, _ := strconv.Atoi(strings.TrimSpace(value))
	return convertToInt

}
