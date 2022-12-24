package main

import "fmt"

func main() {
	fmt.Println("First Class function - Another example")

	firstClassFunction()

	// It is also possible to pass arguments to anonymous function
	func(n string) {
		fmt.Println("Hello", n)
	}("Zeeshan")

	// Another way to define function
	debugFunc := func(text string) string {
		return text
	}
	s := debugFunc
	fmt.Printf("Type of debugFunc: %T\n", debugFunc)
	fmt.Printf("Type of s: %T\n", s)
	fmt.Println(s("Hello world from simple func"))

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
