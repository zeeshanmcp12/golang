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
}
