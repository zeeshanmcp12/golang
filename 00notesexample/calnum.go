package main

import "fmt"

func main() {
	fmt.Println("Calculate Two number!")

	var calNum, num1, num2 int

	fmt.Printf("Select anyone: \n1: Add two numbers \n2: Find the difference \n3: Find Product \n4: Divide two numbers \n")
	fmt.Scanf("%d\n", &calNum)
	fmt.Printf("Thank you for entering %d\n", calNum)

	if calNum == 1 {
		fmt.Printf("Enter two numbers to perform addition (Hit Enter): ")
		fmt.Scanf("%d\n", &num1)
		fmt.Scanf("%d\n", &num2)
		fmt.Println(addTwoNum(num1, num2))

	} else if calNum == 2 {
		fmt.Printf("Enter two numbers to find the difference (Hit Enter): ")
		fmt.Scanf("%d\n", &num1)
		fmt.Scanf("%d\n", &num2)
		fmt.Println(findDifference(num1, num2))

	} else if calNum == 3 {
		fmt.Printf("Enter two numbers to find product (Hit Enter): ")
		fmt.Scanf("%d\n", &num1)
		fmt.Scanf("%d\n", &num2)
		fmt.Println(findProduct(num1, num2))

	} else if calNum == 4 {
		fmt.Printf("Enter two numbers to perform division (Hit Enter): ")
		fmt.Scanf("%d\n", &num1)
		fmt.Scanf("%d\n", &num2)
		fmt.Println(div(num1, num2))

	} else {
		fmt.Println("Invalid Input")
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
