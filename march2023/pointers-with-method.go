package main

import "fmt"

type User struct {
	Name     string
	Username string
}

func main() {
	fmt.Println("Pointers with method")

	// user := User{"Abdullah", "abd@go.dev"}
	user := User{
		Name:     "Abdullah",
		Username: "abd@go.dev",
	}
	fmt.Printf("%#v\n", user)

	user.changeUserName()
	fmt.Printf("%#v\n", user.Username)

}

func (u *User) changeUserName() {
	newUserName := &u.Username
	*newUserName = "abdullah@go.dev"
	fmt.Println(u.Username)

}
