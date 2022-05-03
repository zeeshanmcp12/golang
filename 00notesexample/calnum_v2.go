package main

import "fmt"

var num1, num2 int

func main() {
	fmt.Println("Calculate Two number!")

	var calNum int

	// Calculator logic using loop with code refactoring in function definitions
	for i := 1; i < 5; i++ {
		fmt.Printf("Select anyone: \n1: Add two numbers \n2: Find the difference \n3: Find Product \n4: Divide two numbers \n")
		fmt.Scanf("%d\n", &calNum)
		if calNum == 1 {

			addTwoNum()
			continue

		} else if calNum == 2 {

			findDifference()
			continue

		} else if calNum == 3 {

			findProduct()
			continue

		} else if calNum == 4 {

			div()
			break

		} else {
			fmt.Println("Invalid Input")
			break
		}

	}
}

func addTwoNum() {
	fmt.Printf("Enter two numbers to perform addition (Hit Enter): ")
	fmt.Scanf("%d\n", &num1)
	fmt.Scanf("%d\n", &num2)
	fmt.Printf("Sum of %d and %d is %d\n", num1, num2, num1+num2)

}
func findDifference() {
	fmt.Printf("Enter two numbers to find the difference (Hit Enter): ")
	fmt.Scanf("%d\n", &num1)
	fmt.Scanf("%d\n", &num2)
	fmt.Printf("Difference between %d and %d is %d\n", num1, num2, num1-num2)

}
func findProduct() {
	fmt.Printf("Enter two numbers to find product (Hit Enter): ")
	fmt.Scanf("%d\n", &num1)
	fmt.Scanf("%d\n", &num2)
	fmt.Printf("Product of %d and %d is %d\n", num1, num2, num1*num2)

}
func div() {
	fmt.Printf("Enter two numbers to perform division (Hit Enter): ")
	fmt.Scanf("%d\n", &num1)
	fmt.Scanf("%d\n", &num2)
	fmt.Printf("Division of %d and %d is %d\n", num1, num2, num1/num2)

}
