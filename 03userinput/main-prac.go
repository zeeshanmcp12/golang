package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Taking user input!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please rate your experience: ")

	input, _ := reader.ReadString('\n')

	fmt.Printf("This is what user has input %v\n", input)
	fmt.Printf("Type of user input is %T\n", input)

	fmt.Println("------Conversion started---------")
	strToInt, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("This is what user has input %v\n", strToInt)
	fmt.Printf("Type of strToInt var is %T\n", strToInt)
	fmt.Println("------Conversion completed---------")

}
