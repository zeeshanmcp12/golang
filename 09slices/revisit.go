package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println("Add task")

	var todos = []string{}

	fmt.Print("Add task: ")

	for i := 0; true; i++ {

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		todo_text := strings.TrimSpace(input)

		if todo_text != "done" {
			if todo_text != "" {
				todos = append(todos, todo_text)
				fmt.Print("Add task or type 'done': ")
				continue
			} else {
				fmt.Println("Invalid input! add task or type 'done': ")
			}
		} else if todo_text == "done" {
			fmt.Println("Thank you for adding task: ")
			for j, todo := range todos {
				fmt.Printf("%v task: %v\n", j, todo)
			}
			// fmt.Printf("Select any number to continue: [1]Add Task, [2]Delete Task: ")
			// var index int
			// fmt.Scanf("%v", &index)
			// // if index == 2 {
			// fmt.Print("Select task number to delete: ")
			// var taskIndex int
			// fmt.Scanf("%v", &taskIndex)
			// todos = append(todos[:taskIndex], todos[taskIndex+1:]...)
			// fmt.Println("Task updated: ")
			// for j, todo := range todos {
			// 	fmt.Printf("%v task: %v\n", j, todo)
			// }
			// }
			break
		} else {
			fmt.Println("Invalid input!")
		}

	}
}
