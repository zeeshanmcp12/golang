package main

import (
	"fmt"
)

func main() {
	fmt.Println("Revisiting some concepts")
	var (
		number int
		name   string
	)
	fmt.Print("Enter your name: ")
	fmt.Scanf("%v\n", &name)
	greeter(name)

	fmt.Println()
	fmt.Print("Enter any number: ")
	fmt.Scanf("%v", &number)

	// fmt.Println(number)

	for i := 1; i < 10+1; i++ {
		fmt.Printf("%v x %v = %v\n", number, i, number*i)

	}

}

func greeter(input string) {
	fmt.Printf("Hello %v! Thank you for joining us today.", input)

}
