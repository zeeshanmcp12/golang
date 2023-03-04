package main

import "fmt"

// func main() {
// 	var username string
// 	var age int

// 	fmt.Printf("What is your name: ")
// 	fmt.Scanf("%v\n", &username)
// 	fmt.Printf("How old are you? ")
// 	fmt.Scanf("%v\n", &age)
// 	fmt.Println()

// 	fmt.Printf("Hi %s, Good to know that you are %d years old.", username, age)

// }

func main() {
	var num1, num2, sum int
	fmt.Printf("Enter two numbers to perform addition: ")
	fmt.Scanf("%v%v", &num1, &num2)
	sum = num1 + num2

	fmt.Printf("%d + %d = %d", num1, num2, sum)
}
