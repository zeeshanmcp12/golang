package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")
	// This is one type of loop which is also called while loop
	// execute loop until the condition is true
	number := 0      // initialization
	for number < 8 { // condition
		fmt.Println(number) // Print what was the number at this stage
		number += 1         //post or increment number
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("The Value is %v\n", i)
	}

	// for loop as while loop
	// loop will keep run until the condition is true.
	/*index := 9
	for index < 10 {
		fmt.Println("While loop")
	}*/
}
