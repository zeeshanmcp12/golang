package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Using function in golang!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	input, _ := reader.ReadString('\n')
	f_text := strings.TrimSpace(input)

	greeter(f_text)
}

func greeter(input string) int {
	fmt.Printf("Hello %v, thank you for joining us today!\n", input)
	// var charInName int
	charInName := len(input)
	fmt.Printf("Characters in your name %v", charInName)
	return charInName
	// return charInName

}
