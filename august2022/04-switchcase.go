package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Generating random number!")

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1
	fmt.Println("Dice value: ", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Open up the game!")
	case 2:
		fmt.Println("Move to 2 spots")
	case 3:
		fmt.Println("Move to 3 spots")
	case 4:
		fmt.Println("Move to 4 spots")
	case 5:
		fmt.Println("Move to 5 spots")
	case 6:
		fmt.Println("Roll the dice again and move accordingly!")
	default:
		fmt.Println("Don't know what is it!")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Print("Press enter to continue...")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)

}
