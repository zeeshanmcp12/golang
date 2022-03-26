package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
func main() {
	fmt.Println("----- Todo App V1 -----")

	appItems := []string{"'add task'", "'show task'", "'delete task'"}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Select any 1 from these options, %v: ", appItems)

	input, _ := reader.ReadString('\n')
	f_text := strings.TrimSpace(input)

	switch f_text {
	case "add task":
		fmt.Println("Add new task")
	case "show task":
		fmt.Println("Show your tasks")
	case "delete task":
		fmt.Println("Mark 'done' as complete")
	default:
		fmt.Println("Enter correct value!")

	}

}*/

func main() {
	fmt.Println("Todo App v1")

	// appItems := []string{"[1]add task", "[2]show task"}

	// todos := []string{}

	// fmt.Printf("Type of appItems %T\n", appItems)
	// fmt.Printf("Type of todos %T", todos)
	fmt.Printf("Select any number [1]Add Task, [2]Show Task: ")

	for i := 0; i < 3; i++ {

		f_text := userInputInt()
		if f_text == 1 {

			addTask()
			// continue

		} else if f_text == 2 {
			fmt.Println("Your Todos")
			// showTask(todos)
			// break

		} else {
			fmt.Printf("Invalid Input, Select any number from the list")
		}
	}
}

func addTask() {

	todos := []string{}

	fmt.Printf("Add Task: ")

	for i := 0; true; i++ {
		task := strings.TrimSpace(onlyString())

		if task != "done" {
			if task != "" {
				fmt.Printf("Add task or Enter done: ")
				todos = append(todos, task)
				continue

			} else {
				fmt.Println("Invalid Input, Add task or Enter done: ")

			}
		} else if task == "done" {
			// fmt.Println("Your Todos")
			fmt.Println("Thank you for adding up your tasks!")
			for j, val := range todos {
				fmt.Printf("%v: %v\n", j+1, val)
			}
			fmt.Println()
			break

		}
	}
}

func showTask(tasks []string) {
	// todos := []string{}
	for j, val := range tasks {
		fmt.Printf("%v: %v", j+1, val)
		fmt.Println()
	}
}

func status() {
	fmt.Println("Thank you for adding up your tasks!")

}

func userInputInt() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strToInt(input)

}

func onlyString() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func strToInt(strText string) int {
	input, err := strconv.Atoi(strings.TrimSpace(strText))
	CheckNilErr(err)
	return input
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
