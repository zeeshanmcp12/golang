package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	if 2*5 == 10 {
		fmt.Println("Both are equal")
	} else {
		fmt.Println("There is no equality")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your name: ")
	input, _ := reader.ReadString('\n')
	c_name := strings.TrimSpace(input)
	fmt.Println(c_name)

	if name := c_name; c_name == name {
		fmt.Printf("%v Thank you for login!", c_name)
	} else {
		fmt.Println("Enter correct username")
	}

}
