package main

import "fmt"

func main() {
	fmt.Println("revising golang concepts")

	var tableNumber int

	fmt.Print("Enter any number: ")
	fmt.Scanf("%v\n", &tableNumber)

	fmt.Printf("Thank you for entering %v", tableNumber)

	for i := 0; i < 10+1; i++ {
		fmt.Printf("%v x %v = %v\n", tableNumber, i, tableNumber*i)

	}

}
