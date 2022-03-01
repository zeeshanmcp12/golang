package main

import "fmt"

func main() {
	fmt.Println("defer in golang")

	// This is normal flow of program
	fmt.Println("\n----- This is normal flow -----")
	fmt.Println("1st statement")
	fmt.Println("2nd statement")

	// This is the flow with defer keyword
	/*	fmt.Println("\n----- This is the flow with defer keyword -----")
		defer fmt.Println("1st statement with defer")
		fmt.Println("This will print before defer statement.")*/

	// Combined flow with defer keyword and without it
	fmt.Println("\n----- Combined flow with defer keyword and without it -----")
	defer fmt.Println("Hello --> it will be printed at the end")
	defer fmt.Println("One --> it will be printed at the second last from bottom")
	defer fmt.Println("\nTwo --> it will be printed at third last from bottom")
	fmt.Println("World --> first, this will be printed.")

	fmt.Println("\n----- defer function call started -----")
	deferFunc()

}

// Let's visualize the flow of this func
// Normal flow
// 01234 -> it wont print 5 because condition does not include 5

// Flow with defer keyword
// defer function 43210

func deferFunc() {

	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}

}
