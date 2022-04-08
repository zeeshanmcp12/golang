package main

import "fmt"

func main() {
	fmt.Println("Bitwise Operator in golang!")
	var num1, num2 int = 10, 22
	// 0 0 0 0 1 0 1 0
	// 0 0 0 1 0 1 1 0
	// bitwise AND
	// bitwiseAND := num1 & num2

	// bitwise OR
	// bitwiseOR := num1 | num2

	// bitwise XOR
	bitwiseXOR := num1 ^ num2
	fmt.Println(bitwiseXOR)

	// bitwise leftshift
	bitwiseLeftShift := bitwiseXOR << 1
	fmt.Println(bitwiseLeftShift)

}
