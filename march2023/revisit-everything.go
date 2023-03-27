package main

import "fmt"

func main() {
	fmt.Println("Revisiting and refactoring notes")

	var (
		num1 int = 12
		num2 int = 13
	)

	var num int = 12
	num += 1 // can be written as num = num + 1

	bitwiseAndOperator := num1 & num2
	// fmt.Printf("%b\n%b\n%b\n%d in decimal", num1, num2, bitwiseAndOperator, bitwiseAndOperator)

	bitwiseOROperator := num1 | num2
	fmt.Printf("%b\n%b\n%d in decimal\n%b\n%b\n%d in decimal", num1, num2, bitwiseAndOperator, num1, num2, bitwiseOROperator)

}
