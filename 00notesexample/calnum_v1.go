package main

import "fmt"

func main() {
	fmt.Println("Calculate Two number!")

	var calNum, num1, num2 int

	// Calculator logic using loop
	for i := 1; i < 5; i++ {
		fmt.Printf("Select anyone: \n1: Add two numbers \n2: Find the difference \n3: Find Product \n4: Divide two numbers \n")
		fmt.Scanf("%d\n", &calNum)
		if calNum == 1 {
			fmt.Printf("Enter two numbers to perform addition (Hit Enter): ")
			fmt.Scanf("%d\n", &num1)
			fmt.Scanf("%d\n", &num2)
			fmt.Printf("Sum of %d and %d is %d\n", num1, num2, addTwoNum(num1, num2))
			continue
		} else if calNum == 2 {
			fmt.Printf("Enter two numbers to find the difference (Hit Enter): ")
			fmt.Scanf("%d\n", &num1)
			fmt.Scanf("%d\n", &num2)
			fmt.Printf("Difference between %d and %d is %d\n", num1, num2, findDifference(num1, num2))
			continue

		} else if calNum == 3 {
			fmt.Printf("Enter two numbers to find product (Hit Enter): ")
			fmt.Scanf("%d\n", &num1)
			fmt.Scanf("%d\n", &num2)
			fmt.Printf("Product of %d and %d is %d\n", num1, num2, findProduct(num1, num2))
			continue

		} else if calNum == 4 {
			fmt.Printf("Enter two numbers to perform division (Hit Enter): ")
			fmt.Scanf("%d\n", &num1)
			fmt.Scanf("%d\n", &num2)
			fmt.Printf("Division of %d and %d is %d\n", num1, num2, div(num1, num2))
			break

		} else {
			fmt.Println("Invalid Input")
			break
		}

	}
}

func addTwoNum(num1 int, num2 int) int {
	return num1 + num2
}
func findDifference(num1 int, num2 int) int {
	return num1 - num2
}
func findProduct(num1 int, num2 int) int {
	return num1 * num2
}
func div(num1 int, num2 int) int {
	return num1 / num2
}
