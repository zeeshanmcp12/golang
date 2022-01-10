package main

import (
	"bufio"
	"fmt"
	"os"
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

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Your Name: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello and Welcome!", input)
	fmt.Printf("Type of input variable is %T", input)
}
