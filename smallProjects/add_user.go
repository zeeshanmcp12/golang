package main

import "fmt"

type User struct {
	Name        string
	Age         int
	Email       string
	UserConcent bool
}

func main() {
	fmt.Println("Creating user!")

	platformUser := User{
		"Zeeshan",
		31,
		"abc@abc.com",
		true,
	}

	fmt.Printf("User details: \n%v\n%v\n%v\n%v\n", platformUser.Name, platformUser.Age, platformUser.Email, platformUser.UserConcent)
}
