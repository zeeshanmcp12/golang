package main

import (
	"fmt"
)

// func main() {
// 	fmt.Println("Hello World")

// 	var name string = "Sakina"

// 	for i, j := range name {
// 		fmt.Printf("%v%c\n", i, j)
// 	}
// }

// func main() {
// 	var number int
// 	fmt.Println("Enter any number: ")
// 	fmt.Scanf("%v", &number)

// 	for i := 1; i < 10+1; i++ {
// 		fmt.Printf("%v x %v = %v\n", number, i, number*i)
// 	}
// }

// func main() {
// 	fmt.Println("Enter any text: ")

// 	reader := bufio.NewReader(os.Stdin)

// 	input, _ := reader.ReadString('\n')
// 	fmt.Println("This is what you entered:", input)

// 	fmt.Printf("Type of input is %T\n", input)

// 	strtoInt, err := strconv.Atoi(strings.TrimSpace(input))

// 	fmt.Printf("Now, type of input is %T", strtoInt)

// 	if err != nil {
// 		panic(err)
// 	}
// }

// func main() {
// 	fmt.Println("Array in golang!")

// 	var fruitList [4]string

// 	fruitList[0] = "Apple"
// 	fruitList[1] = "Mango"
// 	fruitList[2] = "Banana"
// 	fruitList[3] = "Orange"

// 	fmt.Println("Elements in an array: ", fruitList)
// 	// fmt.Println("Elements in an array: ", fruitList[1])

// 	// for i := range fruitList {
// 	// 	fmt.Printf("%v -> %v\n", i, fruitList[i])
// 	// }

// 	// var fruitNum int
// 	// fmt.Printf("Type number to select your favorite food: ")
// 	// // fmt.Println()
// 	// fmt.Scanf("%v", &fruitNum)
// 	// fmt.Println(fruitNum)

// 	// for i := 0; i < len(fruitList); i++ {
// 	// 	fmt.Printf("%v -> %v\n", i, fruitList[i])

// 	// }

// 	var anotherList = [3]string{"oneItem", "twoItem", "threeItem"}

// 	for i := range anotherList {
// 		fmt.Printf("%v -> %v\n", i, anotherList[i])
// 	}
// }

// func main() {
// 	fmt.Println("Slice in golang!")

// 	var todos = []string{}

// 	fmt.Printf("Enter your task: ")

// 	for true {

// 		reader := bufio.NewReader(os.Stdin)
// 		input, _ := reader.ReadString('\n')
// 		todoTxt := strings.TrimSpace(input)

// 		if todoTxt != "done" {
// 			if todoTxt != "" {
// 				todos = append(todos, todoTxt)
// 				fmt.Printf("Enter task or type done: ")
// 			} else {
// 				fmt.Printf("Invalid input! Enter task or type done: ")
// 			}
// 			continue
// 		} else if todoTxt == "done" {
// 			fmt.Println("Thank you for adding task. Here is the list: ")
// 			for i, item := range todos {
// 				fmt.Printf("%v -> %v\n", i+1, item)
// 			}
// 		} else {
// 			fmt.Println("Invalid input!, Try again")
// 		}
// 		break
// 	}
// }

// func main() {
// 	fmt.Println("Retrieve item based on index number!")

// 	var list = make([]string, 4)
// 	list[0] = "WFH"
// 	list[1] = "OOO"
// 	list[2] = "BRB"
// 	list[3] = "SL"

// 	fmt.Println("-------- Before append -------- ")
// 	fmt.Println(list)

// 	list = append(list, "PFA", "PFB", "AL")
// 	fmt.Println("-------- After append -------- ")
// 	fmt.Println(list)

// 	var newSlice = append(list[2:6])
// 	fmt.Println("-------- New Slice -------- ")
// 	fmt.Println(newSlice)

// 	fmt.Println("-------- Remove Slice based on index number -------- ")
// 	var indexNum int = 3

// 	var removeItemSlice = append(list[:indexNum], list[indexNum+1:]...)
// 	fmt.Println("-------- Selected item has been removed -------- ")
// 	fmt.Println(removeItemSlice)

// }

// func main() {
// 	fmt.Println("Maps (key value pair) in golang!")
// 	var abbr = make(map[string]string)

// 	abbr["OOO"] = "Out Of Office"
// 	abbr["SL"] = "Sick Leave"
// 	abbr["WFH"] = "Work From Home"
// 	abbr["AL"] = "Annual Leave"

// 	for i, item := range abbr {
// 		fmt.Printf("%v -> %v \n", i, item)
// 	}
// }

// func main() {
// 	fmt.Println("Working with Maps!")

// 	var abbr = make(map[string]string)

// 	abbr["JS"] = "JavaScript"
// 	abbr["RB"] = "Ruby"
// 	abbr["GO"] = "Golang"

// 	for key, item := range abbr {
// 		fmt.Printf("%v -> %v\n", key, item)
// 	}

// 	var strToSelect string

// 	fmt.Printf("Select letter to find out it's abbreviation: ")
// 	fmt.Scanf("%v", &strToSelect)

// 	_, ok := abbr[strToSelect]

// 	fmt.Printf("Type of ok is %T\n", ok)
// 	fmt.Println(ok)

// 	if ok == true {
// 		fmt.Println("Abbreviation found against your provided keyword:", strToSelect)
// 		fmt.Printf("%v stands for %v", strToSelect, abbr[strToSelect])
// 	} else {
// 		fmt.Println("Abbreviation not found. Try again!")
// 	}

// }

// func main() {
// 	fmt.Println("Switch case in golang!")

// 	var name string

// 	fmt.Printf("Enter any name: ")
// 	fmt.Scanf("%v", &name)

// 	fmt.Println(name)

// 	switch name {
// 	case "Abdullah":
// 		fmt.Println("Abdullah is working as Data Scientist")
// 	case "Ahmed":
// 		fmt.Println("A passionate .Net developer")
// 	default:
// 		fmt.Println("No data found!")
// 	}

// }

// func main() {

// 	fmt.Println("Try your luck!")
// 	var ourNum int
// 	fmt.Printf("Enter any number to try your luck: ")
// 	fmt.Scanf("%v", &ourNum)

// 	rand.Seed(time.Now().UnixNano())
// 	number := rand.Intn(10 + 1)

// 	fmt.Println("Your number: ", ourNum)
// 	fmt.Println("Random number: ", number)

// 	switch number {
// 	case ourNum:
// 		fmt.Printf("Congratulations! Your no %v has matched with random no %v ", ourNum, number)
// 	default:
// 		fmt.Printf("No Luck, Try again!")
// 	}

// }

// type User struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Email string `json:"emailAddress"`
// }

// func main() {
// 	fmt.Println("Struct in golang!")

// 	userData := User{"Abdullah", 30, "abc@go.dev"}

// 	fmt.Println(userData)
// 	fmt.Println(userData.Name)
// 	fmt.Printf("%+v\n", userData)
// }

func main() {
	fmt.Println("Function in golang!")
	fmt.Printf("\n ---------- Calling simple function with return values ---------- \n")
	// fmt.Println("Calling addNum simple function with return statement")
	result := addNum(5, 10)
	fmt.Println(result)

	// Simple function call with return statement
	fmt.Println(addNum(10, 20))

	fmt.Printf("\n ---------- Calling function with multiple return values ---------- \n")

	// Simple function call with multiple return
	addition, subtraction := multipleReturnType(10, 5)
	fmt.Printf("Sum is %v , Difference is %v", addition, subtraction)

	fmt.Println()

	fmt.Printf("\n ---------- Calling function with multiple return values ---------- \n")

	var (
		name string
		age  int
	)
	fmt.Printf("Enter your name: ")
	fmt.Scanf("%v \n", &name)
	fmt.Printf("Enter your age: ")
	fmt.Scanf("%v \n", &age)
	fmt.Println()
	userData(name, age)

	fmt.Printf("\n ---------- Calling function with named values ---------- \n")
	// Function with named (return) values
	sum, product := namedValue(2, 5)
	fmt.Printf("Sum of two numbers: %v, Product of two numbers: %v", sum, product)

	// Calling variadic function as described below

	fmt.Printf("\n ---------- Calling Variadic function - Example 1 ---------- \n")
	fmt.Println(variadicFunc()) // in this case, slice will be empty
	fmt.Println(variadicFunc(1, 2))
	fmt.Println(variadicFunc(3, 4, 5))

	// Calling variadic function as described below
	fmt.Printf("\n ---------- Calling Variadic function - Example 2 ---------- \n")
	studentDetails(name, "Physics", "Computer")

	fmt.Printf("\n ---------- Calling function with blank identifier ---------- \n")
	num1, num2 := f()
	fmt.Println(num1, num2)

	onlyNum1, _ := f()
	fmt.Println("Using only 1 value and skipping 2nd using blank identifier: ", onlyNum1)

}

// Simple function with return
func addNum(num1 int, num2 int) int {
	return num1 + num2
}

// Function with multiple return type
func multipleReturnType(num1, num2 int) (int, int) {
	sum := num1 + num2
	diff := num1 - num2
	return sum, diff
}

func userData(name string, age int) (string, int) {
	fmt.Printf("Hey %v, glad you are %v years old!", name, age)
	return name, age
}

func namedValue(num1 int, num2 int) (sum int, product int) {

	sum = num1 + num2
	product = num1 * num2
	return
}

func variadicFunc(numbers ...int) int {
	total := 0

	for _, val := range numbers {
		total += val
	}
	return total
}

// func variadicFunc(numbers ...int) int {
// 	// Declared and initialized a variable with 0
// 	total := 0

// 	// loop through numbers variable (parameter which refer to variadic vaule)
// 	for _, val := range numbers {
// 		// Increament a number to val which is a placeholder for values that will be coming from numbers when we call this function
// 		// After increament, assign the value to 'total' variable
// 		total += val
// 	}
// 	// Return the result (of total)
// 	return total
// }

func studentDetails(name string, subjects ...string) {
	fmt.Printf("Hey %v, Here are your subjects - \n", name)
	for _, sub := range subjects {
		fmt.Printf("%s, ", sub)
	}
}

// Function for blank identifier
func f() (int, int) {
	return 40, 45
}
