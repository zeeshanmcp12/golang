package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
func main() {
	fmt.Println("Reading notes to rephrase the learning!")
	fmt.Println("Taking input from user!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter any text: ")
	input, _ := reader.ReadString('\n')
	converted, _ := strconv.Atoi(strings.TrimSpace(input))
	fmt.Printf("Type of string is %T and value is %v", converted, input)

}*/

/*
func main() {
	fmt.Println("Another main function!")
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(5)
	fmt.Println(randNumber)

	currentTime := time.Now()
	fmt.Println(currentTime.Format("Monday, 2-1-2006"))

	currentDate := time.Date(2006, time.March, 2, 00, 00, 00, 00, time.Local)
	fmt.Println(currentDate)
}*/

func main() {
	fmt.Println("Another function for practice")
	fmt.Printf("Enter your age: ")
	age := userInputinInt()
	fmt.Printf("Enter any number: ")
	luckyNumber := userInputinInt()

	fmt.Printf("Enter your name: ")
	fullname := strings.TrimSpace(onlyString())

	fmt.Printf("Thank you for your input, %v\n%v\n%v", age, luckyNumber, fullname)

}

func userInputinInt() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	// fmt.Printf("Type is %T", strToInt(input))
	return strToInt(input)

}

func onlyString() string {
	reader := bufio.NewReader(os.Stdin)
	stringText, _ := reader.ReadString('\n')
	return stringText

}

func strToInt(anyText string) int {
	atoi, err := strconv.Atoi(strings.TrimSpace(anyText))
	CheckNilErr(err)
	return atoi

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
