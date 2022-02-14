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
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter first number: ")
	firstNum, _ := reader.ReadString('\n')
	fmt.Printf("Enter second number: ")
	secondNum, _ := reader.ReadString('\n')
	fmt.Printf("Select any of these to calculate: 'sum, sub, mul, div': ")
	calculate, _ := reader.ReadString('\n')

	numOne, err := strconv.Atoi(strings.TrimSpace(firstNum))
	numTwo, err := strconv.Atoi(strings.TrimSpace(secondNum))

	var operator string = strings.TrimSpace(calculate)

	if err != nil {
		fmt.Println(err)
	} else {
		// fmt.Printf("Thank you for entering first number %v ", numOne)
		// fmt.Println()
		// fmt.Printf("Thank you for entering second number %v ", numTwo)
		fmt.Println()

		switch operator {
		case "sum":
			fmt.Println("Sum of two numbers: ", numOne+numTwo)
		case "sub":
			fmt.Println("Subtraction of two numbers: ", numOne-numTwo)
		case "mul":
			fmt.Println("Product of two numbers: ", numOne*numTwo)
		case "div":
			fmt.Println("Division of two numbers: ", numOne/numTwo)
		default:
			fmt.Println("Please enter correct value.")

		}
	}
	fmt.Println("\nPress enter to continue...")
	pressAnyKey, _ := reader.ReadString('\n')
	fmt.Println(pressAnyKey)

}
