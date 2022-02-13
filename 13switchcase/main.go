package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch Case in golang!")

	// Write logic similar to Dice game
	// Generate random number as it happens in Dice game
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(6) + 1
	fmt.Printf("Value of Dice is %v", number)
	fmt.Println()

	// Create switch case statement
	switch number {
	case 1: // here 1 means a value we want to execute the case. It could be string here or whatever the case we want to validate
		fmt.Printf("Number is %v, open to Dice", number)
		fmt.Println()
	case 2:
		fmt.Printf("Number is %v, move to 2 spaces", number)
		fmt.Println()
		fallthrough // yet to be understand about this fallthrough in golang
	case 3:
		fmt.Printf("Number is %v, move to 3 spaces", number)
		fmt.Println()
	case 4:
		fmt.Printf("Number is %v, move to 4 spaces", number)
		fmt.Println()
		fallthrough
	case 5:
		fmt.Printf("Number is %v, move to 5 spaces", number)
		fmt.Println()
	case 6:
		fmt.Printf("Number is %v, move to 6 spaces and dice again", number)
		fmt.Println()

	}
}
