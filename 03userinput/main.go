package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var userInput string = "Understanding about how to take user input in golang"
	fmt.Println(userInput)

	reader := bufio.NewReader(os.Stdin)

}
