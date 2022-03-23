package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("New File for practice!")
	fmt.Println()
	fmt.Printf("Enter any number to make table: ")
	tableNum := userInput()

	/*
		for i := 1; i < 10+1; i++ {
			fmt.Printf("%v x %v = %v\n", tableNum, i, tableNum*i)
		}*/

	for i := 1; i < 10+1; i++ {

		fmt.Printf("%v x %v = ", tableNum, i)
		userInput()

		/*
			if tableNum*i == tableNum*i {
				// fmt.Printf("%v x %v = %v\n", tableNum, i, tableNum*i)
				fmt.Printf("%v x %v = %v\n", tableNum, i, userInput())
				continue

			} else if i > 10+1 {
				fmt.Println("Table end!")

			} else {
				fmt.Println("Wrong answer!")
			}*/

	}

}

func userInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strToInt(input)

}

func strToInt(strText string) int {
	input, err := strconv.Atoi(strings.TrimSpace(strText))
	CheckNilErr(err)

	return input
}

func stringInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
