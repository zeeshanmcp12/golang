package main

import "fmt"

func main() {
	fmt.Println("Anonymous function in golang!")

	// One way of declaring anonymous function
	// anonym := func(num1 int, num2 int) int {
	// 	return num1 * num2

	// }
	// fmt.Printf("Type of anonym is %T\n", anonym) // Output: Type of anonym is func(int, int) int
	// fmt.Println(anonym(10, 20))                  // 200

	// Another way of declaring anonymous function
	findProduct := func(num1 int, num2 int) int {
		return num1 * num2
	}(10, 20)

	fmt.Printf("Type of findProduct func is %T\n", findProduct) // Output: Type of findProduct func is int.
	// This time function output result is assigned to variable (findProduct) and since it is an int that's why the type is Int.
	fmt.Println(findProduct) // 200
}
