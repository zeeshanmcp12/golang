package main

import "fmt"

func main() {
	// fmt.Println("Variables!")

	var username string = "Learning Golang!"
	fmt.Println(username)
	fmt.Printf("Type of var username is %T \n", username)
	fmt.Printf("----------------------------------------- \n")

	var isRegistered bool = true
	fmt.Println("Did you Signup today?", isRegistered)
	fmt.Printf("Type of var isRegistered is %T \n", isRegistered)
	fmt.Printf("----------------------------------------- \n")

	var courseAmount uint16 = 256
	fmt.Println(courseAmount)
	fmt.Printf("Type of var courseAmount is %T \n", courseAmount)
	fmt.Printf("----------------------------------------- \n")

	var smallFloat float32 = 255.454678999544555
	fmt.Println(smallFloat)
	fmt.Printf("Type of var smallFloat is %T \n", smallFloat)
	fmt.Printf("----------------------------------------- \n")

	fmt.Printf("Variable declared without mentioning it's type.\n")
	var fullName = "Muhammad Zeeshan"
	fmt.Println(fullName)
	fmt.Printf("Type of var fullName is %T \n", fullName)
	fmt.Printf("----------------------------------------- \n")

	fmt.Printf("Implicit variable declaration.\n")
	text := "This is text"
	fmt.Println(text)
}
