package main

import "fmt"

func main() {
	fmt.Println("Learn & Practice Structs in golang!")
	user := User{"Zeeshan", "acloudtechie@outlook.com", "DevOPS Engineer", 32}
	fmt.Printf("User details are: %v\n", user)
	fmt.Printf("User details are: %+v\n", user)
	fmt.Printf("Username: %v", user.Name)
}

type User struct {
	Name       string
	Email      string
	Profession string
	Age        int
}
