package main

import "fmt"

const anotherVariable = "anotherVariable"

func main() {
	var username string = "zeeshanmcp12"
	fmt.Println(username)
	fmt.Printf("Type of var username is %T \n", username)

	fullName := "Muhammad Zeeshan"
	fmt.Println("My name is ", fullName)

	var isLoggedin = true
	fmt.Printf("Type of var isLoggedin %T \n", isLoggedin)

	fmt.Println(anotherVariable)
	fmt.Printf("Type of const anotherVariable is %T \n", anotherVariable)

}
