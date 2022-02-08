package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("This is about fun exercises!")
	var funnyItems = []string{"Looks good!", "That's really cool", "No way", "Awesome", "Beautiful"}
	// fmt.Println(funnyItems)
	fmt.Println("------------------------------")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a number from 0 to 4 to see something cool: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for entering:", input)
	fmt.Println("------------------------------")

	convertToInt, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println(err)
	} else {
		var index = convertToInt
		fmt.Println(funnyItems[index])
		fmt.Println("------------------------------")

		funnyItems = append(funnyItems[:index], funnyItems[index+1:]...)
		fmt.Println("Remaining items are: ", funnyItems)
		fmt.Println()
	}
	fmt.Printf("Press enter key to continue...")
	pressAnyKey, _ := reader.ReadString('\n')
	fmt.Println(pressAnyKey)
}
