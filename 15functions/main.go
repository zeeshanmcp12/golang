package main

import "fmt"

func main() {
	fmt.Println("Function in golang!")

	// function call
	result := add(5, 5)
	fmt.Printf("Sum of two number is %v ", result)
}

// Write function that can add two numbers
// Function declaration
func add(num1 int, num2 int) int {
	// fmt.Println(num1 + num2)
	return num1 + num2

}
