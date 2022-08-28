package main

import (
	"fmt"
)

func main() {
	fmt.Println("Using function in golang!")

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter your name: ")
	// input, _ := reader.ReadString('\n')
	// f_text := strings.TrimSpace(input)

	// greeter(f_text)

	input := GetInput()
	p := StringData{input}

	fmt.Print("Name: ", p.Name)

}

func greeter(input string) int {
	fmt.Printf("Hello %v, thank you for joining us today!\n", input)
	// var charInName int
	charInName := len(input)
	fmt.Printf("Characters in your name %v", charInName)
	return charInName
	// return charInName

}

type StringData struct {
	Name string
	// Email string
	// Age   int
}

func GetInput() string {
	var input string
	fmt.Print("Enter anything: ")
	fmt.Scanf("%v", &input)
	return input
}
