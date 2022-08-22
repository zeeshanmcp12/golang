package main

import "fmt"

type User struct {
	UserName string
	Name     string
	Email    string
	Age      int
	Concent  bool
}

func main() {
	fmt.Println("Creating/Updating user profile!")

	var (
		username, name, email string
		age                   int
		isAgreed              bool
	)

	fmt.Print("Enter username: ")
	fmt.Scanf("%v\n", &username)

	fmt.Print("Enter name: ")
	fmt.Scanf("%v\n", &name)

	fmt.Print("Enter email: ")
	fmt.Scanf("%v\n", &email)

	fmt.Print("Enter age: ")
	fmt.Scanf("%d\n", &age)

	fmt.Print("Are you agree (true/false): ")
	fmt.Scanf("%t\n", &isAgreed)

	platformUser := User{username, name, email, age, isAgreed}

	fmt.Printf("Type of struct is: %T\n", platformUser)

	fmt.Printf("%+v", platformUser)

}

func (u User) AddUser() {

}

// string, string, string, int, bool
func userDetail() {

}
