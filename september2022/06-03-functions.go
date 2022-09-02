package main

import "fmt"

func main() {
	fmt.Println("using functions in golang!")

	fmt.Println("Sum is: ", adder(10, 10))

	fmt.Println()

	// Calling Variadic function
	fmt.Println("Executing variadic function!")
	fmt.Println("Adding no num: ", variadicFun(0))
	fmt.Println("Adding 2 numbers: ", variadicFun(2, 4))
	fmt.Println("Adding 3 numbers: ", variadicFun(2, 3, 4))

	fmt.Println()

	// Anonymous function
	fmt.Println("Executing anonymous function!")
	result := func(num1, num2 int) int {
		sum := num1 + num2
		return sum
	}(10, 20)

	fmt.Printf("Type of result: %T\n", result)
	fmt.Println("Sum is", result)

	fmt.Println()
	fmt.Println("Executing anonymous function v2!")
	anonymFunc := func(num1, num2 int) int {
		sum := num1 + num2
		return sum
	}
	fmt.Printf("Type of func: %T\n", anonymFunc)
	fmt.Println("Sum is: ", anonymFunc(20, 30))

}

func adder(num1, num2 int) int {
	sum := num1 + num2
	return sum
}

func variadicFun(number ...int) int {
	sum := 0

	for _, value := range number {
		sum += value
	}
	return sum
}
