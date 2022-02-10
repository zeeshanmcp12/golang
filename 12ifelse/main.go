package main

import "fmt"

func main() {
	fmt.Println("If/Else in golang!")

	loginCount := 10
	var result string

	if loginCount > 15 {
		result = "Regular User"
		fmt.Println(result)
	} else if loginCount < 15 {
		result = "Visit the site to win vouchers!"
		fmt.Println(result)
	} else {
		result = "Exactly 10 login count"
		fmt.Println(result)
	}

	if isLoggedIn := true; !isLoggedIn {
		fmt.Println("Thank you for login!")
	} else {
		fmt.Println("Enter correct username")
	}
}
