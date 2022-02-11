package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("New file for practice!")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter your name: ")
	input, _ := reader.ReadString('\n')

	c_userName := strings.TrimSpace(input)
	fmt.Println(c_userName)

	userDb := []string{}
	userDb = append(userDb, c_userName)
	fmt.Printf("Value in userDb is %v", userDb)

}
