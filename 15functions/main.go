package main

import "fmt"

func main() {
	fmt.Println("Function in golang!")

	// function call
	// result := add(5, 5)
	// fmt.Printf("Sum of two number is %v ", result)

	// fmt.Println(data(32, "Zeeshan"))
	result := sub(14, 7)
	fmt.Printf("14 - 7 is %v", result)
}

// Write function that can add two numbers
// Function declaration
/*func add(num1 int, num2 int) int {
	// fmt.Println(num1 + num2)
	return num1 + num2

}*/
// Mention data type only once because of same type of parameter
/*func add(num1, num2 int) int {
	// fmt.Println(num1 + num2)
	return num1 + num2
}*/

// Defining two data types together
/*func data(age int, name string) (int, string) {
	return age, name

}*/

// labeling data type of return
func sub(num1, num2 int) (lbl1 int) {
	lbl1 = num1 - num2
	return
}
