package main

import "fmt"

func main() {
	fmt.Println("Hello from April 2023")

	var username string

	fmt.Printf("Enter your name: ")

	fmt.Scanf("%v", &username)

	fmt.Println("Hi", username, "Thank you for joining us!")
}
