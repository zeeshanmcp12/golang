package main

import "fmt"

const LoginName string = "zeeshanmcp12" // Public, so it is available across the files of program.

func main() {
	// fmt.Println("Variables!")

	var username string = "Learning Golang!"
	fmt.Println(username)
	fmt.Printf("Type of var username is %T \n", username)
	fmt.Printf("----------------------------------------- \n")
	// This % sign(before letter 'T') is a placeholder.

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

	// Default value and some aliases
	var intVariable int
	fmt.Println(intVariable)
	// This variable is just declared and not initialized and golang will not add garbage value into it but print 0 because it is integer as we described it's type.
	fmt.Printf("Type of var intVariable is %T \n", intVariable)
	fmt.Printf("----------------------------------------- \n")

	var stringVariable string
	fmt.Println(stringVariable)
	// This variable is just declared and not initialized. golang will not add garbage value into it but print nothing if the variable is of type string.
	fmt.Printf("Type of var stringVariable is %T \n", stringVariable)
	fmt.Printf("----------------------------------------- \n")

	// Implicit type or an implicit way of declaring variables
	fmt.Printf("Variable declared without mentioning it's type.\n")
	var fullName = "Muhammad Zeeshan"
	fmt.Println(fullName)
	fmt.Printf("Type of var fullName is %T \n", fullName)
	fmt.Printf("----------------------------------------- \n")
	// In this case lexer will identify whether it is of type string or else. But, we cannot reinitialize this variable with different type i.e. int or float etc. It will give error saying "already declared this varibale to string"
	// fullName = 3 -> this is not allowed.

	// No var style of declaring variable.
	text := "This is text"
	fmt.Println(text)
	fmt.Printf("No var style of declaring variable. \n")
	fmt.Printf("----------------------------------------- \n")
	// No var keyword here. We just used ":="
	// Note: Inside of any method we can use this volorous operator (:=) to declare a variable but outside of method it is not allowed. for example global declaration or outside of main function.

	// for const declaration, please go to top.

	fmt.Println(LoginName)
	fmt.Printf("The type of LoginName variable is %T \n", LoginName)
}
