package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")
	/*
		// This is one type of loop which is also called while loop
		// execute loop until the condition is true
		number := 0      // initialization
		for number < 8 { // condition
			fmt.Println(number) // Print what was the number at this stage
			number += 1         //post or increment number
		}*/
	/*
		for i := 0; i < 10; i++ {
			fmt.Printf("The Value is %v\n", i)
		}*/
	sum := 0
	for i := 1; i < 5; i++ {
		fmt.Printf("Sum inside loop %v\n", sum)
		sum += i
	}
	fmt.Printf("Sum outside of loop %v\n", sum)
	/*	sum := 0
		sum += 1
		fmt.Println(sum)
		sum += 1
		fmt.Println(sum)
		sum += 1
		fmt.Println(sum)
		sum += 1
		fmt.Println(sum)
		sum += 1
		fmt.Println(sum)
		sum += 1
		fmt.Println(sum)*/

	// for loop as while loop
	// loop will keep run until the condition is true.
	/*index := 9
	for index < 10 {
		fmt.Println("While loop")
	}*/
}
