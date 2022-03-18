package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Practice from 6 to 10")
	fmt.Printf("Enter any number: ")
	anything := userInputtoInt()
	fmt.Printf("Type is %T: ", anything)
	fmt.Println(anything)

	fmt.Printf("Enter any text: ")
	anyText := onlyString()
	stringText := strings.TrimSpace(anyText)
	fmt.Printf("Type is %T: ", stringText)

}

func userInputtoInt() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strtoint(input)

}

func onlyString() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input

}

func strtoint(stringText string) int {
	strText, err := strconv.Atoi(strings.TrimSpace(stringText))
	CheckNilErr(err)
	return strText

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}

}
