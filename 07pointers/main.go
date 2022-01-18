package main

import "fmt"

func main() {
	fmt.Println("This is about Pointers in golang!")

	// var ptr *int
	// fmt.Println("The value of ptr variable is", ptr)

	myNumber := 26
	var ptr = &myNumber
	fmt.Println("Value of ptr is", ptr)
	fmt.Println("Added 1 into myNumber:", myNumber+1)

}
