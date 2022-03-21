package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Logic for Calculator!")
	fmt.Printf("Enter 1st number: ")
	oneNum := userInput()

	fmt.Printf("Enter 2nd number: ")
	twoNum := userInput()

	fmt.Printf("Select anyone to calculate, sum, sub, mul, div:")
	operator := strings.TrimSpace(onlyString())

	switch operator {
	case "sum":
		fmt.Println("Sum of two number is: ", oneNum+twoNum)
	case "sub":
		fmt.Println("Subtraction of two number is: ", oneNum-twoNum)
	case "mul":
		fmt.Println("Product of two number is: ", oneNum*twoNum)
	case "div":
		fmt.Println("Divsion of two number is: ", oneNum/twoNum)
	default:
		fmt.Println("Invalid Input!")

	}

}

func userInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strToInt(input)

}

func strToInt(strText string) int {
	converted, _ := strconv.Atoi(strings.TrimSpace(strText))
	return converted
}

func onlyString() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
