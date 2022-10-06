package main

import "fmt"

func main() {
	fmt.Println("Resume learning go in October 2022!")

	var (
		age  int
		name string
	)

	fmt.Print("Enter your name: ")
	fmt.Scanf("%v\n", &name)
	// fmt.Println()

	fmt.Print("Enter your age: ")
	fmt.Scanf("%v\n", &age)
	fmt.Println()

	fmt.Printf("Hello %v, good to know that you are %v years old.", name, age)
}
