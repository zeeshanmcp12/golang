package main

import "fmt"

func main() {
	fmt.Println("functions in golang!")

	// fmt.Println()

	var numOne, numTwo int
	fmt.Printf("Enter two numbers: ")
	fmt.Scanf("%v,%v", &numOne, &numTwo)
	fmt.Printf("Sum of %v and %v is %v", numOne, numTwo, proAdder(numOne, numTwo))
	fmt.Println()

	// add()

}

// func add() {
// 	var numOne, numTwo int
// 	fmt.Println("------- Simple addition -------")

// 	fmt.Printf("Enter two numbers (separated with comma) to calculate the sum: ")
// 	fmt.Scanf("%v,%v", &numOne, &numTwo)
// 	fmt.Printf("Sum of %v and %v is %v", numOne, numTwo, numOne+numTwo)
// }

func proAdder(numOne, numTwo int) int {

	return numOne + numTwo
}
