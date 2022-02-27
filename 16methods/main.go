package main

import "fmt"

func main() {
	fmt.Println("Methods in golang!")

	user := User{"Zeeshan", true, 32, "abc@abc.com"}
	fmt.Println("User is", user.Name)
	fmt.Println("Old email address is: ", user.Email)
	user.UserStatus()
	user.NewEmail() // here it will print manipulated email address, because it's only a copy being passed on here and not an original value which is abc@abc.com
	// Reason is on line 17
	fmt.Println("Old email address is: ", user.Email) // here it will still print old email address.
}

// Manipulate struct with method
// Whenever we pass on these objects (or structs) it actually passes on a copy, here comes the concept of pointers in golang
// So to pass the original object, we should be passing up the reference of it (or a pointer to that).
func (u User) NewEmail() {
	u.Email = "acloudtechie@outlook.com"
	fmt.Println("New email is: ", u.Email)
}

func (u User) UserStatus() {
	fmt.Println("Is user logged in:", u.isActive)
}

type User struct {
	Name     string
	isActive bool
	Age      int
	Email    string
}
