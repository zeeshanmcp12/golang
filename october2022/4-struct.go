package main

import "fmt"

func main() {
	fmt.Println("Data structure in golang using struct!")

	var (
		name, email string
		age         int
	)

	fmt.Print("Enter your name: ")
	fmt.Scanf("%v\n", &name)

	fmt.Print("Enter your email: ")
	fmt.Scanf("%v\n", &email)

	fmt.Print("Enter your age: ")
	fmt.Scanf("%v\n", &age)

	// userData := User{"Zeeshan", "abc@go.dev", 30}

	userData := User{name, email, age}

	fmt.Printf("Hey %v, your email address is: %v\n", userData.Name, userData.Email)
	fmt.Printf("Glad to know that you are %v years old.\n", userData.Age)

	fmt.Printf("Data in details: %+v", userData)
	// fmt.Println(userData.Name)

}

type User struct {
	Name  string
	Email string
	Age   int
}
