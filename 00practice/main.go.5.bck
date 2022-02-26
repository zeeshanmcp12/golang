package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Function in golang!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter 1st number: ")
	inputOne, _ := reader.ReadString('\n')
	fmt.Printf("Enter 2nd number: ")
	inputTwo, _ := reader.ReadString('\n')

	num1, _ := strconv.Atoi(strings.TrimSpace(inputOne))
	num2, _ := strconv.Atoi(strings.TrimSpace(inputTwo))

	add := add(num1, num2)
	fmt.Printf("Sum of two number is: %v", add)

	subtract := sub(num1, num2)
	fmt.Println()
	fmt.Printf("Subtraction of two number is: %v", subtract)

	multiply := multiply(num1, num2)
	fmt.Println()
	fmt.Printf("Product of two number is: %v", multiply)

	divide := div(num1, num2)
	fmt.Println()
	fmt.Printf("Division of two number is: %v", divide)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}
func sub(num1 int, num2 int) int {
	return num1 - num2
}
func multiply(num1 int, num2 int) int {
	return num1 * num2
}
func div(num1 int, num2 int) int {
	return num1 / num2
}
