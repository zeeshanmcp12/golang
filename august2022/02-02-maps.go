package main

import "fmt"

func main() {
	fmt.Println("Map literal!")

	var myMap = map[string]int{
		"phoneNumber":  +9212345678,
		"homeNumber":   +9232154789,
		"officeNumber": +92513465678,
	}

	fmt.Println(myMap)
}
