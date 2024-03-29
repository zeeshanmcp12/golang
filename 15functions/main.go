package main

import "fmt"

func main() {
	fmt.Println("Function in golang!")

	// function call
	// result := add(5, 5)
	// fmt.Printf("Sum of two number is %v ", result)

	// fmt.Println(data(32, "Zeeshan"))
	// result := sub(14, 7)
	// fmt.Printf("14 - 7 is %v", result)

	// return value by using variadic function
	result := unlimitedValue(3, 2, 1)
	fmt.Println(result)

	// writing function inside function
	output := func() {
		fmt.Println("function call from inside of function")
	}
	output()

	user := User{"Zeeshan", 32, true, "abc@abc.com"}

	fmt.Println("Name is: ", user.Name)
	fmt.Println("Age is: ", user.Age)
	fmt.Println("Email is: ", user.Email)
}

// Write function that can add two numbers
// Function declaration
/*func add(num1 int, num2 int) int {
	// fmt.Println(num1 + num2)
	return num1 + num2

}*/
// Mention data type only once because of same type of parameter
/*func add(num1, num2 int) int {
	// fmt.Println(num1 + num2)
	return num1 + num2
}*/

// Defining two data types together
/*func data(age int, name string) (int, string) {
	return age, name

}*/

// label our return values
// lbl1 -> is the label for our return value
/*func sub(num1, num2 int) (lbl1 int) {
	lbl1 = num1 - num2
	return
}*/
func unlimitedValue(values ...int) int {
	total := 0

	for _, val := range values {
		total += val // this (val) value will be added every time to total and total is being upgraded.

	}
	// now, return the value
	return total
}

type User struct {
	Name   string
	Age    int
	Status bool
	Email  string
}
