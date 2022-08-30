package main

import "fmt"

func main() {
	fmt.Println("Handling user data!")

	// Initialize structure with values
	// platformUser := User{"Abdullah", "abdullah@go.dev", 30, true}

	// fmt.Println(platformUser)
	// fmt.Printf("User details: %+v\n", platformUser)
	// fmt.Printf("Username: %v", platformUser.Name)

	var (
		username, email string
		age             int
		isVerified      bool
	)

	fmt.Print("Enter username: ")
	fmt.Scanf("%v\n", &username)

	fmt.Print("Enter email: ")
	fmt.Scanf("%v\n", &email)

	fmt.Print("How old are you: ")
	fmt.Scanf("%d\n", &age)

	fmt.Print("Do you agree: ")
	fmt.Scanf("%t\n", &isVerified)

	platformUser := User{username, email, age, isVerified}
	fmt.Printf("Thank you for registeration %v\n", platformUser.Name)
	fmt.Println("------------------------")
	fmt.Printf("Your Profile:\n%+v\n%+v\n%d\n%t\n", platformUser.Name, platformUser.Email, platformUser.Age, platformUser.isLoggedIn)
	fmt.Println("------------------------")

	fmt.Println(platformUser.UpdateEmail())

}

// Define structure (or struct)
type User struct {
	Name       string
	Email      string
	Age        int
	isLoggedIn bool
}

func (u User) UpdateEmail() string {

	updatedEmail := "updatedemail@go.dev"

	return updatedEmail

}
