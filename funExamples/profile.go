package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
1- Take input for user details -> Done
2- Convert input to string/int -> Done
3- Define structure -> Done
4- Utilize elements from structure -> Done
5- Validate if user exist -> not done yet
*/

func main() {
	fmt.Println("\tCreate Profile!")
	fmt.Println()

	// Take input for user details
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your Name: ")
	username, _ := reader.ReadString('\n')
	fmt.Printf("Enter Email Address: ")
	emailAddress, _ := reader.ReadString('\n')
	fmt.Printf("Enter your Age: ")
	age, _ := reader.ReadString('\n')

	// Convert input to string/int
	// Age converted to Int because it was string before
	convertedAge, _ := strconv.Atoi(strings.TrimSpace(age))

	// Trim space from string because of ReadString function above
	convertedUsername := strings.TrimSpace(username)
	convertedEmail := strings.TrimSpace(emailAddress)

	// Utilize elements from structure
	userProfile := User{convertedUsername, convertedEmail, convertedAge}
	fmt.Println()
	fmt.Printf("----------Your Profile----------\n")
	fmt.Printf("Name: %v\nEmail: %v\nAge: %v\n\n", userProfile.Name, userProfile.Email, userProfile.Age)
	fmt.Println("--------------------------------")

	fmt.Println("Press enter to continue...")
	pressAnyKey, _ := reader.ReadString('\n')
	fmt.Println(pressAnyKey)
}

// Define Structure
type User struct {
	Name  string
	Email string
	Age   int
}
