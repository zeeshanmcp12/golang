package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Slice in golang!")

	var todos = []string{}

	fmt.Println("Add your task: ")

	for i := 0; true; i++ {

		// fmt.Println("Add your task: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		todo_t := strings.TrimSpace(input)

		if todo_t != "done" {
			if todo_t != "" {
				todos = append(todos, todo_t)
				fmt.Println("Type 'done' or Add todo: ")
				continue

			} else {
				fmt.Println("Invalid input, type 'done' or Add todo: ")
			}
		} else if todo_t == "done" {
			fmt.Println("Thank you for adding tasks: ")
			for j, todo := range todos {
				fmt.Printf("%v task: %v\n", j+1, todo)
			}
			break

		} else {
			fmt.Println("Invalid input!")
		}

	}
}
