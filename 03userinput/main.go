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
	fmt.Println("Please rate how much you understood the concept:")
	// notes are in notes.md file for above syntax

	// comma ok syntax || err err
	input, _ := reader.ReadString('\n')
	//input -> whatever an input will be given
	// _ -> (underscore) means, if any mistake/error occurs during input so this is something try catch.
	// reader -> is a variable we declared above
	// .ReadString -> is a method which will wait (or read) for string we gonna input.
	// ('\n') -> represents new line so reader.ReadString will read for string till we press "Enter".
	fmt.Println("You gave us:", input)
	fmt.Printf("Type of input var is %T", input)

}
