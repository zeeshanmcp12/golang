package main

import "fmt"

func main() {
	fmt.Println("Function in golang!")
	result := add(5, 6)
	fmt.Printf("Sum of two number is: %v", result)
}

func add(num1 int, num2 int) int {
	return num1 + num2
}
