package main

import "fmt"

type User struct {
	UserName string
	Name     string
	Email    string
	Age      int
	Concent  bool
}

var userInput string

func main() {
	fmt.Println("Creating/Updating user profile!")

	var (
		username, name, email string
		age                   int
		isAgreed              bool
		updateEmailConsent    string
	)

	fmt.Print("Enter username: ")
	fmt.Scanf("%v\n", &username)

	fmt.Print("Enter name: ")
	fmt.Scanf("%v\n", &name)

	fmt.Print("Enter email: ")
	fmt.Scanf("%v\n", &email)

	fmt.Print("Enter age: ")
	fmt.Scanf("%d\n", &age)

	fmt.Print("Do you agree (true/false): ")
	fmt.Scanf("%t\n", &isAgreed)

	platformUser := User{username, name, email, age, isAgreed}

	fmt.Printf("%+v\n", platformUser)

	fmt.Print("Enter 'yes' if you want to update the email: ")
	fmt.Scanf("%v\n", &updateEmailConsent)

	if updateEmailConsent == "yes" {
		fmt.Print("Enter email: ")
		fmt.Scanf("%v", &userInput)

		platformUser.UpdateUser()
	} else if updateEmailConsent == "no" {
		fmt.Println("Thank you for registeration!")
	} else {
		fmt.Printf("Invalid input!")
	}

}

func (u User) UpdateUser() {
	// fmt.Print("Enter email: ")
	// fmt.Scanf("%v", &userInput)
	u.Email = userInput

	fmt.Printf("Email has been updated: %v", u.Email)

}
