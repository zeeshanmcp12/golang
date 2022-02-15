package main

import "fmt"

func main() {
	fmt.Printf("----------Loops in Golang----------\n")

	days := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
		// fmt.Printf("Value is %v\n", days[d])

	}
}
