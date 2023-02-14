package main

import "fmt"

type User struct {
	name  string
	email string
}

func (u *User) updateEmail(email string) {
	u.email = email

}

func main() {
	fmt.Println("Method in golang using struct")
	fmt.Printf("Empty struct %+v\n", User{})

	user := User{
		name:  "Abdullah",
		email: "abc@go.dev",
	}
	fmt.Printf("Struct with values -> %+v\n", user)

	user.updateEmail("emailChange@go.dev")
	fmt.Printf("Struct with values -> %+v\n", user)

}
