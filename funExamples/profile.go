package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
1- Take input for user details
2- Convert input to string/int
3- Define structure
4- Utilize elements from structure
5- Validate if user exist

*/
func main() {
	fmt.Println("Create Profile!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter username: ")
	username, _ := reader.ReadString('\n')
	fmt.Printf("Enter email address: ")
	emailAddress, _ := reader.ReadString('\n')
	fmt.Printf("Enter Age: ")
	age, _ := reader.ReadString('\n')

	convertedAge, _ := strconv.Atoi(strings.TrimSpace(age))
	// convertedUsername, _ := strconv.Atoi(strings.TrimSpace(username))
	convertedUsername := strings.TrimSpace(username)
	convertedEmail := strings.TrimSpace(emailAddress)

	fmt.Printf("Types of variable %T, %T, %T", username, emailAddress, age)
	fmt.Println()
	fmt.Printf("Types of variable %T %T %T", convertedUsername, convertedEmail, convertedAge)
	fmt.Println()
	fmt.Printf("Your Profile: %v, %v, %v", convertedUsername, convertedEmail, convertedAge)

}
