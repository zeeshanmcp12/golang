package main

import "fmt"

func main() {
	fmt.Println("Returning values in golang!")

	// fmt.Println("Sum of two numbers:", addNum(5, 4))

	result := func(num1, num2 int) int {
		sum := num1 * num2
		return sum
	}
	fmt.Printf("Type of function is %T\n", result)
	fmt.Println("Sum is", result(10, 20))

	annonymusFun := func(num1, num2 int) int {
		sum := num1 * num2
		return sum
	}(20, 30)
	fmt.Printf("Type of function is %T\n", annonymusFun)
	fmt.Println("Product is", annonymusFun)

}
