package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Todo app functions using High order function!")

	tasks := []string{}

	fmt.Printf("Add tasks: ")

	fmt.Printf("Enter \n 1- Add Task\n 2- Show Task")

	addTask(tasks...)

	// for true {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	input, _ := reader.ReadString('\n')
	// 	f_txt := strings.TrimSpace(input)

	// 	if f_txt != "done" {
	// 		if f_txt != "" {
	// 			tasks = append(tasks, f_txt)
	// 			fmt.Printf("Type done or Press Enter: ")
	// 		} else {
	// 			fmt.Printf("Invalid input! Type done or Press Enter: ")
	// 			continue
	// 		}
	// 	} else if f_txt == "done" {
	// 		fmt.Printf("Thank you for adding tasks:\n")
	// 		for i, val := range tasks {
	// 			fmt.Printf("%v -> %v\n", i, val)
	// 		}
	// 		break
	// 	} else {
	// 		fmt.Printf("Invalid input! Try again.")
	// 	}
	// }
}

func addTask(tasks ...string) {
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
