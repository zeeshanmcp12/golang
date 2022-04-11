package main

import "fmt"

func main() {
	fmt.Println("Switch Case in golang!")

	// var number int = 100
	var num1, num2 int = 10, 20

	switch {
	case num1+num2 == 30:
		fmt.Println("Greater than 30")
	case num1+num2 >= 30:
		fmt.Println("Greater than equal to 30")
		// fallthrough
	// case 2:
	// fmt.Println("What to do next")
	default:
		fmt.Println("No matched!")
	}

	// Tricky question in kodekloud lab.
	// If below code is not in editor then we need to find out either there will be an error or any result.
	// This will through an error because switch contains "int" value while cases contains bool and it is not allowed to convert untyped bool to int. That's why it will give error.
	var a, b = 100, 5
	switch a {
	case a/b == 10:
		fmt.Println("10")
	case a/b == 20:
		fmt.Println("20")
	case a/b == 10:
		fmt.Println("30")
	default:
		fmt.Println("default")
	}
}
