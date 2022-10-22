package main

import "fmt"

func main() {
	fmt.Println("functions in golang!")

	var numOne, numTwo int
	fmt.Printf("Enter two numbers: ")
	fmt.Scanf("%v,%v", &numOne, &numTwo)

	// Calling proAdder function here
	fmt.Printf("Sum of %v and %v is %v\n", numOne, numTwo, proAdder(numOne, numTwo))
	fmt.Println("----------------------")

	// Writing and calling anonymous function
	anonymousfuncv1 := func(val1, val2 int) int {
		return val1 + val2
	}(3, 4)

	fmt.Printf("Type of anonymousfuncv1 func %T\n", anonymousfuncv1)
	fmt.Println("Result of anonymousfuncv1: ", anonymousfuncv1)
	fmt.Println("----------------------")

	// Writing and calling anonymous function v2
	anonymousfuncv2 := func(num1, num2 int) int {
		return num1 + num2
	}

	fmt.Printf("Type of anonymousfuncv2 is %T\n", anonymousfuncv2)
	fmt.Printf("Result of aanonymousfuncv2 is: %v\n", anonymousfuncv2(2, 3))
	fmt.Println("----------------------")

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
