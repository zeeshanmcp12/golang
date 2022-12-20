package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Todo app functions using High order function!")

	var query int

	// fmt.Printf("Enter \n 1- Show Task\n 2- Add Task: ")
	fmt.Printf("Enter \n 1- Addition\n 2- Subtraction: ")
	fmt.Scanf("%d\n", &query)

	getFunction(query)

	// addTask(tasks...)

}

func addTask(tasks ...string) {

	// tasks := []string{}

	fmt.Printf("Add tasks: ")

	for true {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		f_txt := strings.TrimSpace(input)

		if f_txt != "done" {
			if f_txt != "" {
				tasks = append(tasks, f_txt)
				fmt.Printf("Type done or Press Enter: ")
			} else {
				fmt.Printf("Invalid input! Type done or Press Enter: ")
				continue
			}
		} else if f_txt == "done" {
			fmt.Printf("Thank you for adding tasks:\n")
			for i, val := range tasks {
				fmt.Printf("%v -> %v\n", i, val)
			}
			break
		} else {
			fmt.Printf("Invalid input! Try again.")
		}
	}
}

func showTask(tasks ...string) {
	for i, val := range tasks {
		fmt.Printf("%v -> %v\n", i, val)
	}
}

func addition(num1, num2 int) {
	fmt.Println(num1 + num2)
	// return result
}

func subtraction(num1, num2 int) {
	fmt.Println(num1 - num2)
	// return result
}

func getFunction(query int) func(int) {

	query_to_func := map[int]func(int){
		1: addition,
		2: subtraction,
	}
	return query_to_func[query]

}

// func getFunction(query int) func(r float64) float64 {

// 	query_to_func := map[int]func(r float64) float64{
// 		1: calcArea,
// 		2: calcPerimeter,
// 		3: calcDiameter,
// 	}
// 	return query_to_func[query]
// }
