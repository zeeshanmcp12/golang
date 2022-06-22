package main

import "fmt"

func main() {
	fmt.Println("Understanding Zero Values in Golang!")

	var username string
	var age int
	var isLoggedIn bool

	fmt.Printf("username is %s\n", username)
	fmt.Printf("age is %d\n", age)
	fmt.Printf("Did you login today? %t\n", isLoggedIn)
}
