package main

import (
	"fmt"
	"log"
	"time"
)

// func main() {
// 	fmt.Println("Bitwise Operator in Golang!")

// 	var (
// 		num1 int = 12
// 		num2 int = 13
// 	)

// 	var num int = 12
// 	num += 1 // can be written as num = num + 1

// 	bitwiseAndOperator := num1 & num2
// 	// fmt.Printf("%b\n%b\n%b\n%d in decimal", num1, num2, bitwiseAndOperator, bitwiseAndOperator)

// 	bitwiseOROperator := num1 | num2
// 	fmt.Printf("%b\n%b\n%d in decimal\n%b\n%b\n%d in decimal", num1, num2, bitwiseAndOperator, num1, num2, bitwiseOROperator)

// }

// func main() {

// 	fmt.Println("Switch-Case in Golang!")

// 	switch os := runtime.GOOS; os {
// 	case "darwin":
// 		fmt.Println("It's OSX")
// 	case "Linux":
// 		fmt.Println("Linux yeah")
// 	default:
// 		fmt.Println("No! Its Windows")

// 	}
// }

func vanquish() {
	fmt.Println("vanquishing the dragon")
	start := time.Now()
	time.Sleep(2 * time.Second)
	fmt.Println("Dragon Vanquished!")
	elapsed := time.Since(start)
	// fmt.Printf("Execution: %s", elapsed)
	log.Printf("Execution: %s", elapsed)
}

func main() {
	// vanquish the dragon
	vanquish()
}
