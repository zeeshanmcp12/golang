package main

import "fmt"

func main() {
	fmt.Println("Handling user data!")

	// Initialize structure with values
	var personInfo = User{
		Name:       "Zeeshan",
		Email:      "abc@abc.com",
		Age:        32,
		isLoggedIn: true,
	}

	fmt.Printf("%+v\n", personInfo)

	var person2Info = User{
		Name: "Hussain",
	}

	fmt.Println("Peron 2: ", person2Info)

}

// Define structure (or struct)
type User struct {
	Name       string
	Email      string
	Age        int
	isLoggedIn bool
}
