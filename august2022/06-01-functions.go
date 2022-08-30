package main

import "fmt"

func main() {
	fmt.Println("Returning values in golang!")

	// fmt.Println("Sum of two numbers:", addNum(5, 4))

	// Returning multiple values
	fmt.Print(addNum(5, 6))

	fmt.Println()
	// Variadic function
	fmt.Println(variadicFunc())
	fmt.Println(variadicFunc(10))
	fmt.Println("Sum of these numbers:", variadicFunc(5, 6, 7))

	fmt.Println()
	// Without blank identifier
	result, sum := blankIdnt()
	fmt.Println("Number1:", result, "Number2:", sum)

	// With blank identifier
	result1, _ := blankIdnt()
	fmt.Println("Only Number1", result1)

}

// Returning multiple values
func addNum(num1, num2 int) (string, int) {
	sum := num1 + num2
	// stringReturned := "Sum of two numbers: "
	return "Sum of two numbers: ", sum

}

// Variadic functions
// Here in function signature, argument (numbers) will be slice
func variadicFunc(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum

}

// Blank identifier
// It is used to simply ignore the values that are returned by function
func blankIdnt() (int, int) {
	return 40, 54

}
