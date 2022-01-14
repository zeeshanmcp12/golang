package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// presentTime := time.Now()
	// fmt.Println(presentTime)
	// fmt.Println(presentTime.Format("2006-01-02 Monday"))

	createdDate := time.Date(2021, time.December, 03, 07, 00, 00, 00, time.Local)
	fmt.Println(createdDate)
	fmt.Println("----------------------------------------------")
	fmt.Println("CNIC Issue Date:", createdDate.Format("02-01-2006 15:04 Monday"))
	fmt.Println("----------------------------------------------")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please enter your DOB - YY,MM,DD Day:")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for your input: ", input)
	convertInputToInt, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println(err)
	} else {
		dateOfBirth := time.Date(convertInputToInt, 00, 00, 00, 00, time.Local)
		fmt.Println(dateOfBirth.Format("02-01-2006 Monday"))
	}

}
