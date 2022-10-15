package main

import "fmt"

func main() {
	fmt.Println("Revisiting structs in golang!")

	var (
		name        string
		age         int
		isConfirmed bool
	)

	fmt.Print("Enter your name: ")
	fmt.Scanf("%v\n", &name)

	fmt.Print("Enter your age: ")
	fmt.Scanf("%v\n", &age)

	fmt.Print("Do you agree? (true/false) ")
	fmt.Scanf("%t\n", &isConfirmed)

	userData := User{name, age, isConfirmed}

	// fmt.Println("Hi", userData.Name, "glad to know you are", userData.Age, "years old.")

	if userData.isLoggedIn {
		fmt.Println("It is from if", userData.isLoggedIn)
	} else {
		fmt.Println("It is from else", userData.isLoggedIn)
	}

	// fmt.Printf("Hi %q, glad to know you are %v years old.", userData.Name, userData.Age)

	// fmt.Printf("Type of struct is %T", userData)
}

type User struct {
	Name       string
	Age        int
	isLoggedIn bool
}
