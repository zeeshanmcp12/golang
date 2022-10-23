package main

import "fmt"

func main() {
	fmt.Println("Method in golang!")
	fmt.Println("When we use 'Struct' in function, it is called method.")

	userData := User{"Zeeshan", 30, "zeeshan@go.dev", true}

	fmt.Println(userData)
	userData.GetStatus()
	userData.UpdateEmail()
	fmt.Println("Previous email was:", userData.Email)
}

type User struct {
	Name       string
	Age        int
	Email      string
	IsLoggedIn bool
}

func (u User) GetStatus() {
	fmt.Println("Is User active?", u.IsLoggedIn)
}

func (u User) UpdateEmail() {
	u.Email = "zeeshan@acloudtechie.com"
	fmt.Println("Record has been updated:", u.Email)
}
