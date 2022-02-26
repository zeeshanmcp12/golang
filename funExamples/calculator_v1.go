package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("----------Basic Caluclator----------")

	fmt.Printf("Enter first number: ")
	firstNum := userInput()
	fmt.Printf("Enter second number: ")
	secondNum := userInput()

	fmt.Printf("Select any of these to calculate: 'sum, sub, mul, div': ")

	calculate := onlyInput()

	var operator string = strings.TrimSpace(calculate)

	switch operator {
	case "sum":
		fmt.Println()
		fmt.Printf("Sum of %v and %v is : %v ", firstNum, secondNum, firstNum+secondNum)

	case "sub":
		fmt.Println()
		fmt.Printf("Remaining from %v and %v is : %v ", firstNum, secondNum, firstNum-secondNum)

	case "mul":
		fmt.Println()
		fmt.Printf("Product of %v and %v is : %v ", firstNum, secondNum, firstNum*secondNum)

	case "div":
		fmt.Println()
		fmt.Printf("Division of %v and %v is : %v ", firstNum, secondNum, firstNum/secondNum)

	default:
		fmt.Println("Please enter correct value.")

	}
	fmt.Println("\nPress enter to continue...")
	pressAnyKey := onlyInput()
	fmt.Println(pressAnyKey)

}

func onlyInput() string {
	reader := bufio.NewReader(os.Stdin)
	stringText, _ := reader.ReadString('\n')
	return stringText

}

func userInput() int {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strtoint(text)
}

func strtoint(value string) int {
	convertToInt, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		fmt.Println(err)
	}
	return convertToInt
}
