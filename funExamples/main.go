package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("This is about fun excersices!")
	var funnyItems = []string{"Looks good!", "You're cool", "no way", "Awesome", "Beautiful"}
	fmt.Println(funnyItems)
	// fmt.Println(len(funnyItems))

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter number from 1 to 4 to see something cool: ")
	input, _ := reader.ReadString('\n')
	fmt.Printf(input)
	convertToInt, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println(err)
	} else {
		var index = convertToInt
		fmt.Println(funnyItems[index])
		funnyItems = append(funnyItems[:index], funnyItems[index+1:]...)
		fmt.Println("Remaining items are: ", funnyItems)
	}
}
