package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Match number and win the game!")

	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10)

	// fmt.Println("Random number:", randNum)

	var ourNum int
	fmt.Print("Enter any number (1-10): ")
	fmt.Scanf("%d", &ourNum)
	fmt.Println("Our number:", ourNum)

	switch randNum {
	case ourNum:
		fmt.Printf("Congratulations! Your number %d is matched with random number %d", ourNum, randNum)
	default:
		fmt.Printf("No match found! Your number: %d, random number: %d", ourNum, randNum)

	}

}
