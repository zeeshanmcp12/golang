package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// const anotherVariable = "anotherVariable"

func main() {
	// var username string = "zeeshanmcp12"
	// fmt.Println(username)
	// fmt.Printf("Type of var username is %T \n", username)

	// fullName := "Muhammad Zeeshan"
	// fmt.Println("My name is ", fullName)

	// var isLoggedin = true
	// fmt.Printf("Type of var isLoggedin %T \n", isLoggedin)

	// fmt.Println(anotherVariable)
	// fmt.Printf("Type of const anotherVariable is %T \n", anotherVariable)
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Please enter your name: ")
	// input, _ := reader.ReadString('\n')
	// fmt.Println("Thank you for joining us,", input)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("Enter Your Name: ")
	// input, _ := reader.ReadString('\n')
	// fmt.Println("Hello and Welcome!", input)
	// fmt.Printf("Type of input variable is %T", input)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter any number: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("This number you entered:", input)

	convertInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Conversion completed! ", convertInput+1)
		fmt.Printf("Type of convertInput var is %T", convertInput)
	}
}
