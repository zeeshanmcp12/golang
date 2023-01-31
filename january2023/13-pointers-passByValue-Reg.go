package main

import "fmt"

func main() {
	fmt.Println("Pass by reference to function")
	// var numbers = []int{10, 20, 30}
	// fmt.Printf("Type of numbers %T", numbers)

	var number int = 100

	var ptr = &number

	*ptr = 200

	fmt.Printf("Address of number is %v\n", &number)
	fmt.Printf("Address of ptr is %v and Value at the address of ptr %v\n", &ptr, *ptr)
	fmt.Printf("value changed at number to %v\n", number)

	*ptr = 300
	fmt.Printf("value changed at number to %v\n", number)

	arr := []int{10, 20, 30}
	fmt.Println("Values stored in array:", arr)

}
