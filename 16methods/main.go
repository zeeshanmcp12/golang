package main

import "fmt"

func main() {
	fmt.Println("Methods in golang!")

	user := User{"Zeeshan", true, 32}
	fmt.Println("User is", user.Name)
	user.UserStatus()
}

type User struct {
	Name     string
	isActive bool
	Age      int
}

func (u User) UserStatus() {
	fmt.Println("Is user logged in:", u.isActive)
}
