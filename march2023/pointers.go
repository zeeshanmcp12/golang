package main

import "fmt"

func main() {
	fmt.Println("Pointers in golang!")

	number := 42
	fmt.Printf("Address of number variable %v\n", &number)

	newNumber := &number
	fmt.Printf("Value of newNumber %v is the address of number\n", newNumber)
	fmt.Printf("Value of *newNumber %v is the value at that address \n", *newNumber)

	*newNumber = 50
	fmt.Printf("Address of *newNumber %v\n", newNumber)
	fmt.Printf("Value at the address of *newNumber %v\n", number)

}
