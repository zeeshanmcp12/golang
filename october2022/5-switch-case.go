package main

import "fmt"

func main() {
	fmt.Println("Switch case in golang!")

	var userInput int

	fmt.Print("Enter any number: ")
	fmt.Scanf("%v", &userInput)
	fmt.Printf("Thank you for input: %v\n", userInput)

	switch userInput {
	case 5:
		fmt.Printf("Case 5 matched with your number %v", userInput)
	case 3:
		fmt.Printf("Case 3 matched with your number %v", userInput)
	default:
		fmt.Println("Number doesn't match so default statement printed!")
	}
}
