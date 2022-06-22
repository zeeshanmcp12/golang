package main

import "fmt"

func main() {
	fmt.Println("Taking user input using Scanf function!")

	var weight float32
	var height float32

	fmt.Print("Enter your weight: ")
	fmt.Scanf("\n", &weight)

	fmt.Print("Enter your height: ")
	fmt.Scanf("\n", &height)

	fmt.Printf("Your weight is %.2f and height is %.2f", weight, height)
}
