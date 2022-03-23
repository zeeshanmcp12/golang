package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var userDb = []string{"Zeeshan", "Ali", "Hussain"}

func main() {
	fmt.Println("Login to system!")

	// Take input for user details
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your Name: ")
	username, _ := reader.ReadString('\n')

	// Trim space from string because of ReadString function above
	convertedUsername := strings.TrimSpace(username)

	// Utilize elements from structure
	userProfile := User{convertedUsername}

	if userProfile.Name == userDb[0] {
		fmt.Printf("Thank you for login, %v", userProfile.Name)
		fmt.Println()
	} else if userProfile.Name != userDb[0] {
		fmt.Println("User does not exist")
	} else {
		fmt.Println("Enter correct details")
	}
	/*
		fmt.Println()
		fmt.Printf("----------Your Profile----------\n")
		fmt.Printf("Name: %v\nEmail: %v\n", userProfile.Name, userProfile.Email)
		fmt.Println("--------------------------------")*/

	fmt.Println("Press enter to continue...")
	pressAnyKey, _ := reader.ReadString('\n')
	fmt.Println(pressAnyKey)
}

// Define Structure
type User struct {
	Name string
}
