package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in golang!")

	var userInput int

	fmt.Print("Enter any number: ")
	fmt.Scanf("%v", &userInput)

	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10)

	fmt.Printf("Your number: %v\n", userInput)
	fmt.Printf("Random number: %v\n", randNum)

	switch userInput {
	case randNum:
		fmt.Printf("Congratulations! Random no: %v matched with your number %v", randNum, userInput)
	default:
		fmt.Println("Number doesn't match so default statement printed!")
	}
}
