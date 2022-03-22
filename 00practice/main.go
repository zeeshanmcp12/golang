package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Struct
/*
func main() {
	fmt.Println("Reading notes from 11 to 15")
	userProfile := Profile{"Zeeshan", 32, "acloudtechie@outlook.com"}
	fmt.Println(userProfile)
}*/

/*
// if/else
func main() {
	fmt.Println("Reading notes about if/else!")


		userProfile := Profile{"Abdullah", 32, "abc@abc.com"}

		if userProfile.Name == "Zeeshan" && userProfile.Email == "abc@abc.com" {
			fmt.Printf("Thank you %s \n", userProfile.Name)
		} else {
			fmt.Printf("%s, Please enter correct information!\n", userProfile.Name)
		}

		if num := 2; num < 3 {
			fmt.Println(false)
		} else if num == 3 {
			fmt.Println(true)

		} else {
			fmt.Println("else statement")
		}

}
*/
// SwitchCase
/*
func main() {
	fmt.Println("Switchcase and if/else together")
	fmt.Printf("Enter 1st number: ")
	firstNumber := userInputToInt()

	fmt.Printf("Enter 2nd number: ")
	secondNumber := userInputToInt()

	fmt.Printf("Select anyone to calculate, sum, sub, mul, div: ")
	operator := strings.TrimSpace(onlyString())

	if operator == "sum" || operator == "sub" || operator == "mul" || operator == "div" {
		switch operator {
		case "sum":
			fmt.Println("Sum of two number is", firstNumber+secondNumber)
		case "sub":
			fmt.Println("Subtraction of two number is", firstNumber-secondNumber)
		case "mul":
			fmt.Println("Product of two number is", firstNumber*secondNumber)
		case "div":
			fmt.Println("Division of two number is", firstNumber/secondNumber)
		}

	} else {
		fmt.Println("Invalid Input!")
	}

}
*/

// Loops
func main() {
	fmt.Println("Loops in golang")

	/*
		output := funcForLoop(4)
		fmt.Println(output)

		result := addNumbers(4, funcForLoop(4))
		fmt.Println(result)

		for i := funcForLoop(5); i < 10; i++ {
			fmt.Println("Hello world")

		}*/

	days := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	// loop to print index along with the value
	for i := 0; i < len(days); i++ {
		fmt.Printf("Day is %v at %v location\n", days[i], i)
	}

	// Similar method to above
	fmt.Println()
	for i, day := range days {
		fmt.Printf("Day is %v at %v location\n", day, i)

	}
}

func addNumbers(num1, num2 int) int {
	result := num1 + num2
	return result
}

func funcForLoop(num int) int {
	return num
}

type Profile struct {
	Name  string
	Age   int
	Email string
}

func userInputToInt() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return stringToInt(input)

}

func stringToInt(strText string) int {
	input, _ := strconv.Atoi(strings.TrimSpace(strText))
	return input
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
