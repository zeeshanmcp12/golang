package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Let's try our luck!")

	var number int

	fmt.Print("Enter any number: ")
	fmt.Scanf("%v\n", &number)
	fmt.Printf("You entered %v\n", number)

	// anotherNumber := rand.Seed(time.Now().UnixNano())

	switch number {
	case 1, 10:
		rand.Seed(time.Now().UnixNano())
		fmt.Println(rand.Intn(number))
	case 2:
		fmt.Println("Move dice to 2 moves!")
	default:
		fmt.Println("Invalid!")

	}
	fmt.Println("Thank you!")
}
