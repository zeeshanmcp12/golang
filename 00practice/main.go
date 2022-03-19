package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
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

}*/

func main() {
	fmt.Println("Practice 6 to 10")
	fmt.Println("-----------------------")
	fmt.Printf("Enter date of your birth: ")
	date := userInputtoInt()

	fmt.Printf("Enter month of your birth: ")
	month := userInputtoInt()

	fmt.Printf("Enter year of your birth: ")
	year := userInputtoInt()

	birthDate := time.Date(year, time.Month(month), date, 00, 00, 00, 00, time.Local)
	fmt.Printf("Happy Birthday on %v\n", birthDate.Format("02 Jan, 2006"))
	fmt.Printf("It was %v!\n", birthDate.Format("Monday"))

}

func userInputtoInt() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strToInt(input)
}

func strToInt(strText string) int {
	converted, err := strconv.Atoi(strings.TrimSpace(strText))
	CheckNilErr(err)
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
