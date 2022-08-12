package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Win the game by matching your number with randomly generated")

	var yourNum int

	fmt.Print("Enter your number(1-10): ")
	fmt.Scanf("%d\n", &yourNum)
	fmt.Println("Your number: ", yourNum)

	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10)
	fmt.Println("Random number: ", randNum)

	switch randNum {
	case yourNum:
		fmt.Println("Congratulations! Number matched.")
	default:
		fmt.Println("Try again!")

	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press enter to continue...")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)

	// for true {
	// 	if randNum == yourNum {
	// 		reader := bufio.NewReader(os.Stdin)
	// 		fmt.Print("Press enter to continue...")
	// 		input, _ := reader.ReadString('\n')
	// 		fmt.Println(input)
	// 		break

	// 		// fmt.Println("Number Matched!")
	// 	} else {
	// 		// fmt.Println("Continue....")
	// 		fmt.Print("Enter your number(1-10): ")
	// 		fmt.Scanf("%d\n", &yourNum)
	// 		fmt.Println("Your number: ", yourNum)

	// 		rand.Seed(time.Now().UnixNano())
	// 		randNum := rand.Intn(10)
	// 		fmt.Println("Random number: ", randNum)
	// 		continue
	// 	}
	// }

}
