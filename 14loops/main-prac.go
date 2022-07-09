package main

import "fmt"

func main() {
	fmt.Println("Loops in golang!")

	var tabNumber int

	fmt.Print("Enter any number: ")
	fmt.Scanf("%v\n", &tabNumber)

	fmt.Println("Thank you for entering", tabNumber)

	for i := 1; i < 10+1; i++ {
		fmt.Printf("%v x %v = %v\n", tabNumber, i, tabNumber*i)

	}
}
