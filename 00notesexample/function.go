package main

import "fmt"

func main() {
	fmt.Println("functions in golang!")
	fmt.Println("Sum of 3 and 2 is", addNum(3, 2))
}

func addNum(num1 int, num2 int) int {
	return num1 + num2
}
