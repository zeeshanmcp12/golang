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

	fmt.Printf("Enter \n 1- Show Task\n 2- Add Task: ")
	// fmt.Printf("Enter \n 1- Addition\n 2- Subtraction: ")
	fmt.Scanf("%d\n", &query)
	// tasks := []string{}

	// getFunction(query)
	// printResult(tasks, getFunction(query))
	printResult(getFunction(query))

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

func printResult(calculateFunc func(item ...string)) {
	result := calculateFunc
	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Thank you!")
}

func getFunction(query int) func(item ...string) {

	query_to_func := map[int]func(item ...string){
		1: addTask,
		2: showTask,
	}
	return query_to_func[query]

}
