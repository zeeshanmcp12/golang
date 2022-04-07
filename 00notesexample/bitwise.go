package main

import "fmt"

func main() {
	fmt.Println("Bitwise Operator in golang!")
	var num1, num2 int = 10, 22
	// 0 0 0 0 1 0 1 0
	// 0 0 0 1 0 1 1 0
	// bitwiseAND := num1 & num2

	bitwiseOR := num1 | num2
	fmt.Println(bitwiseOR)

}
