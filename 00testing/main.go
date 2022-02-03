package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// const anotherVariable = "anotherVariable"

func main() {
	// var username string = "zeeshanmcp12"
	// fmt.Println(username)
	// fmt.Printf("Type of var username is %T \n", username)

	// fullName := "Muhammad Zeeshan"
	// fmt.Println("My name is ", fullName)

	// var isLoggedin = true
	// fmt.Printf("Type of var isLoggedin %T \n", isLoggedin)

	// fmt.Println(anotherVariable)
	// fmt.Printf("Type of const anotherVariable is %T \n", anotherVariable)
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Please enter your name: ")
	// input, _ := reader.ReadString('\n')
	// fmt.Println("Thank you for joining us,", input)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("Enter Your Name: ")
	// input, _ := reader.ReadString('\n')
	// fmt.Println("Hello and Welcome!", input)
	// fmt.Printf("Type of input variable is %T", input)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("Please enter any number: ")
	// input, _ := reader.ReadString('\n')
	// fmt.Println("This number you entered:", input)

	// convertInput, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Conversion completed! ", convertInput+1)
	// 	fmt.Printf("Type of convertInput var is %T", convertInput)
	// }

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("Please enter anything: ")
	// input, _ := reader.ReadString('\n')
	// fmt.Printf(input)
	// fmt.Printf("Type of input var is %T", input)
	// fmt.Println()

	// converted, err := strconv.Atoi(strings.TrimSpace(input))
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("Type of input is now %T", converted)
	// }

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("Enter any number: ")
	// input, _ := reader.ReadString('\n')
	// // fmt.Printf(input)

	// convertedToInt, err := strconv.Atoi(strings.TrimSpace(input))
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	// fmt.Println("1 added in given input:", convertedToInt+1)
	// 	fmt.Printf("Type of convertedToInt is %T", convertedToInt)
	// 	fmt.Println()
	// 	rand.Seed(int64(convertedToInt))
	// 	fmt.Println(rand.Intn(10))
	// }

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("Enter any number: ")
	// input, _ := reader.ReadString('\n')
	// fmt.Println("Thank you for your input: ", input)
	// convertedtoInt, err := strconv.Atoi(strings.TrimSpace(input))
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("String converted to Int and added a num:", convertedtoInt+1)
	// }
	// fmt.Printf("Type of input var is %T", convertedtoInt)

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(10))

	// randomNum, _ := rand.Int(rand.Reader, big.NewInt(10))
	// fmt.Println(randomNum)
	// presentTime := time.Now()
	// fmt.Println(presentTime)
	// fmt.Println(presentTime.Format("2006-01-02 Monday"))

	// currentTime := time.Now()
	// fmt.Println(currentTime)
	// formattedTime := currentTime.Format("2006-01-02 Monday")
	// fmt.Println(formattedTime)

	// currentTime := time.Now()
	// fmt.Println(currentTime)
	// fmt.Println(currentTime.Format("2006-01-02 15:04:05 Monday"))

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(10))

	// randomN, _ := rand.Int(rand.Reader, big.NewInt(10))
	// fmt.Println(randomN)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("Enter Year: ")
	// year, _ := reader.ReadString('\n')
	// // reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("Enter Month: ")
	// month, _ := reader.ReadString('\n')
	// // reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("Enter Date: ")
	// date, _ := reader.ReadString('\n')

	// Year, err := strconv.Atoi(strings.TrimSpace(year))
	// Month, err := strconv.Atoi(strings.TrimSpace(month))
	// Date, err := strconv.Atoi(strings.TrimSpace(date))

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("-------------- DOB ---------------")
	// 	dateOfBirth := time.Date(Year, time.Month(Month), Date, 00, 00, 00, 00, time.Local)
	// 	fmt.Println(dateOfBirth.Format("02 Jan 2006"))
	// 	fmt.Println("-------- Day of Birth ----------")
	// 	fmt.Println(dateOfBirth.Format("It was Monday!"))
	// }
	// fmt.Println("Press Enter to continue...")
	// pressedKey, _ := reader.ReadString('\n')
	// fmt.Printf(pressedKey)

	// var ptr *int
	// fmt.Println(ptr)

	// number := 12
	// var ptr = &number
	// fmt.Println("Value of ptr var is", ptr)
	// fmt.Println("Value of ptr var is", *ptr)

	// *ptr = *ptr + 2
	// fmt.Println(number)

	// =================================== Practice ==============================================
	// fmt.Println("Hello World!")
	/* 	reader := bufio.NewReader(os.Stdin)
	   	fmt.Printf("Enter any number: ")
	   	input, _ := reader.ReadString('\n')
	   	fmt.Println(input)
	   	convertedNum, err := strconv.Atoi(strings.TrimSpace(input))
	   	if err != nil {
	   		fmt.Println(err)
	   	} else {
	   		fmt.Println(convertedNum)
	   		fmt.Printf("Type of convertedNum is %T", convertedNum)
	   	} */
	// rand.Seed(time.Now().Unix())
	// fmt.Println(rand.Intn(10))

	// presentTime := time.Now()
	// fmt.Println(presentTime.Format("Monday, 02 Jan 2006"))
	// createdAt := time.Date(2006, time.January, 02, 00, 00, 00, 00, time.Local)
	// fmt.Println(createdAt.Format())

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("Enter Year of Birth: ")
	// year, _ := reader.ReadString('\n')

	// fmt.Printf("Enter Month of Birth: ")
	// month, _ := reader.ReadString('\n')

	// fmt.Printf("Enter Date of Birth: ")
	// date, _ := reader.ReadString('\n')

	// yearConvtoInt, err := strconv.Atoi(strings.TrimSpace(year))
	// monthConvtoInt, err := strconv.Atoi(strings.TrimSpace(month))
	// dateConvtoInt, err := strconv.Atoi(strings.TrimSpace(date))

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	createdAt := time.Date(yearConvtoInt, time.Month(monthConvtoInt), dateConvtoInt, 00, 00, 00, 00, time.Local)
	// 	fmt.Println(createdAt)
	// 	fmt.Println(createdAt.Format("Monday, 02 Jan 2006"))
	// }
	// fmt.Println("Practicing about Arrays!")
	// var list [4]string
	// list[0] = "Apple"
	// list[1] = "Mango"
	// list[2] = "Peach"
	// list[3] = "Orange"
	// fmt.Println("List contains these fruits:", list)
	// fmt.Println("Items in list:", len(list))
	// fmt.Println("Count of 1st item in array:", len(list[0]))

	// var iniList = [3]string{"Potatoe", "Mushroom", "Lemon"}
	// fmt.Println(iniList)

	fmt.Println("Starting again with golang")
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("Enter any number: ")
	// input, _ := reader.ReadString('\n')
	// fmt.Printf("Type of input is %T", input)
	// fmt.Println("")
	// converted, err := strconv.Atoi(strings.TrimSpace(input))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("Type of input is %T", converted)
	// fmt.Println("")

	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5))

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter year of your Birth: ")
	year, _ := reader.ReadString('\n')
	fmt.Printf("Enter month of your Birth: ")
	month, _ := reader.ReadString('\n')
	fmt.Printf("Enter day of your Birth: ")
	day, _ := reader.ReadString('\n')

	convertYear, _ := strconv.Atoi(strings.TrimSpace(year))
	convertMonth, _ := strconv.Atoi(strings.TrimSpace(month))
	convertDay, _ := strconv.Atoi(strings.TrimSpace(day))

	dob := time.Date(convertYear, time.Month(convertMonth), convertDay, 00, 00, 00, 00, time.Local)
	convertedDate := dob.Format("Monday, 02-01-2006")
	fmt.Println(convertedDate)

}
