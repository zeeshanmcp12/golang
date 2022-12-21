package main

import "fmt"

func main() {
	fmt.Println("First Class function - Another example")

	firstClassFunction()

}

func firstClassFunction() {

	addition := func(n1, n2 int) int {
		return n1 + n2
	} //(10, 20)
	// fmt.Printf("Type of addition %T\n", addition)
	// fmt.Println(addition(10, 20))

	multiplication := func(n1, n2 int) int {
		return n1 * n2
	}

	funcArr := []func(int, int) int{addition, multiplication}

	for _, function := range funcArr {
		fmt.Println(function(5, 2))
	}

}
