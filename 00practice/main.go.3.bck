package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("New file for practice!")
	/*	reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter your name: ")
		input, _ := reader.ReadString('\n')

		c_userName := strings.TrimSpace(input)
		fmt.Println(c_userName)

		userDb := []string{}
		userDb = append(userDb, c_userName)
		fmt.Printf("Value in userDb is %v", userDb)*/
	/*if num := 10; num < 10 {
		fmt.Println("Number is less than 10")
	} else if num > 10 {
		fmt.Println("Number is NOT less than 10")
	} else {
		fmt.Println("Number is equal")
	}*/

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter any number to compare: ")
	input, _ := reader.ReadString('\n')
	c_number, _ := strconv.Atoi(strings.TrimSpace(input))
	fmt.Printf("Thank you for entering: %v", c_number)
	fmt.Println()
	if num := 5; c_number < num {
		fmt.Printf("Your number %v is less than %v defined in if/else", c_number, num)
	} else if c_number > num {
		fmt.Printf("Your number %v is NOT than %v defined in if/else", c_number, num)
	} else {
		fmt.Printf("Both %v and %v Number is equal", c_number, num)
	}
}
