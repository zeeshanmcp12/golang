package main

import (
	"fmt"
)

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
	/*for i := range days {
		fmt.Println(days[i])
	}*/
	// when we are using this i here, it returns an index and not the actual value.

	// for-each kind of loop in golang ----- it's kind of important to understand ------
	// If we want index and it's value together then we can use below format.
	// first value that we return is i which can also be called as index and we can also return the value itself.
	// We can call 2nd value whatever we like, in this case we used day
	// we can initiliaze this day var with range.
	// this range will range through over the days (slice.)
	// When we want to return two value then we have to keep in mind that first will be index and second will be it's value. But this is only when we are supposed to use it with range keyword.
	/*for index, day := range days {
		fmt.Printf("Index is %v and Value is %v\n", index+1, day)
	}*/

	// ------------------- Incorrect logic to get working day ---------------
	/*index := 0
	weekDayStart := days[2]
	// strConv, _ := strconv.Atoi(weekDayStart)
	for index < len(days) {
		if days[2] == weekDayStart {
			fmt.Println(weekDayStart)
			fmt.Println("Stop! It's Monday")
			break
		}
		fmt.Println("Outside of if/else executed.")
	}*/
	fmt.Println("-----------------------------------------")

	for i, day := range days {
		fmt.Printf("Index is %v and Value is %v\n", i, day)
		if i == 2 {
			fmt.Println(i)
			fmt.Printf("It's %v\n", day)
			break
		}
		i++
	}
	/*
		if index < len(days) {
			fmt.Println(len(days))
			fmt.Printf("Value is %v", days[3])

		} else {
			fmt.Println("Else executed.")
		}*/
}
