package main

import "fmt"

func main() {
	fmt.Println("Method in golang!")
	fmt.Println("When we use 'Struct' in function, it is called method.")

	userData := User{"Zeeshan", 30, true}

	fmt.Println(userData)
	userData.GetStatus()
}

type User struct {
	Name       string
	Age        int
	IsLoggedIn bool
}

func (u User) GetStatus() {
	fmt.Println("Is User active?", u.IsLoggedIn)
}
