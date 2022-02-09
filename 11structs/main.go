package main

import "fmt"

func main() {
	fmt.Println("Learn & Practice Structs in golang!")

	// Let's utilize our structure
	newUser := User{"Zeeshan", 30, "acloudtechie@outlook.com", "DevOPS Engineer"}
	fmt.Println(newUser)

	// %v stands for value
	// %+v This will print more details like Key:Value for example Name:Zeeshan
	fmt.Printf("User details are: %+v\n", newUser)

	fmt.Printf("Username: %v\nEmail: %v", newUser.Name, newUser.Email)
}

// Define structure
type User struct {
	Name       string
	Age        int
	Email      string
	Profession string
}
