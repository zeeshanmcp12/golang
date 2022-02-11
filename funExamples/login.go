package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Login to system!")

	// Take input for user details
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your Name: ")
	username, _ := reader.ReadString('\n')
	fmt.Printf("Enter Email Address: ")
	emailAddress, _ := reader.ReadString('\n')
	// fmt.Printf("Enter your Age: ")
	// age, _ := reader.ReadString('\n')

	// Convert input to string/int
	// Age converted to Int because it was string before
	// convertedAge, _ := strconv.Atoi(strings.TrimSpace(age))

	// Trim space from string because of ReadString function above
	convertedUsername := strings.TrimSpace(username)
	convertedEmail := strings.TrimSpace(emailAddress)

	// Utilize elements from structure
	userProfile := User{convertedUsername, convertedEmail}

	if userProfile.Name == userProfile.Name {
		fmt.Printf("Thank you for login, %v", userProfile.Name)
		fmt.Println()
	} else if userProfile.Name != userProfile.Name {
		fmt.Printf("Your provided %v does not match with our system.", userProfile.Name)

	} else {

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
	Name  string
	Email string
}
