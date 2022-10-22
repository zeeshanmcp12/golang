package main

import "fmt"

func main() {
	fmt.Println("Looping through data structure")

	days := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	fmt.Println(days)

	// Simple loop
	for i := 0; i < len(days); i++ {
		fmt.Printf("This is %v\n", days[i])

	}
	fmt.Println("------------- loop through the range keyword -------------")
	// Range through the loop
	for j, val := range days {
		fmt.Printf("Index -> %v, Value -> %v\n", j, val)
	}

	fmt.Println("------------- loop with continue -------------")
	number := 1

	for number < 5 {

		if number == 3 {
			number++
			continue

		}

		fmt.Println(number)
		number++
	}

	fmt.Println("------------- loop with break -------------")
	secondNumber := 1

	for secondNumber < 10 {
		if secondNumber == 5 {
			break
		}
		fmt.Println(secondNumber)
		secondNumber++

	}

}
