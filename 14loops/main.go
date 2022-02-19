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
	/*	sum := 0
		for i := 1; i < 5; i++ {
			fmt.Printf("Sum inside loop %v\n", sum)
			sum += i
		}
		fmt.Printf("Sum outside of loop %v\n", sum)
	*/
	// for loop as while loop
	// loop will keep run until the condition is true.
	/*number := 0
	for number < 5 {
		fmt.Println(number)
		number++
	}*/

	days := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Firday"}
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// If we want to loop through the entire range then the prefered method is to create for loop with range keyword.
	for i := range days {
		fmt.Println(days[i])
	}
	// when we are using this i here, it returns an index and not the actual value.

	// If we want index and it's value together then we can use below format.
	for index, day := range days {
		fmt.Printf("Index is %v and Value is %v\n", index+1, day)
	}
}
