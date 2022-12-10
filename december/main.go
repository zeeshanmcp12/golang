package main

import "fmt"

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

func main() {
	fmt.Println("Switch case in golang!")

	var name string

	fmt.Printf("Enter any name: ")
	fmt.Scanf("%v", &name)

	fmt.Println(name)

	switch name {
	case "Abdullah":
		fmt.Println("Abdullah is working as Data Scientist")
	case "Ahmed":
		fmt.Println("A passionate .Net developer")
	default:
		fmt.Println("No data found!")
	}

}

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
