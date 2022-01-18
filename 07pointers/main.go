package main

import "fmt"

func main() {
	fmt.Println("This is about Pointers in golang!")

	// var ptr *int
	// fmt.Println("The value of ptr variable is", ptr)

	myNumber := 26                       // Initializing var and assigned value to it.
	var ptr = &myNumber                  // Re-initializing var with already initialized variable. But this time referencing with & sign
	fmt.Println("Value of ptr is", ptr)  // It will print the direct location of memory (or memory address) of variable
	fmt.Println("Value of ptr is", *ptr) // The value inside this *ptr (pointer) will be 26.
	// When I put *ptr, it means I want to see what's inside that pointer?
	// How I was able to fill that?
	// I was able to fill that by adding a reference to &myNumber.
	// *ptr means it's not memory address. It's an actual value inside of it which is 26 in our case.
}
