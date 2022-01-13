package main

import (
	"fmt"
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
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("2006-01-02 Monday"))

}
