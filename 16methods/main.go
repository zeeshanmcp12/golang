package main

import "fmt"

func main() {
	fmt.Println("Methods in golang!")

	user := User{"Zeeshan", true, 32}
	fmt.Println("User is", user.Name)
}

type User struct {
	Name     string
	isActive bool
	Age      int
}
