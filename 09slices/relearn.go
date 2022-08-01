package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Revisiting slice concept in golang!")

	tasks := []string{}

	fmt.Print("Add tasks: ")

	for i := 0; true; i++ {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		todo_text := strings.TrimSpace(input)

		if todo_text != "done" {
			if todo_text != "" {
				tasks = append(tasks, todo_text)
				fmt.Printf("Add task or type done: ")
				continue
			} else {
				fmt.Printf("Invalid input, add task or type done: ")
			}
		} else if todo_text == "done" {
			fmt.Println("Thank you for adding tasks: ")
			for j, todo := range tasks {
				fmt.Printf("%v task: %v\n", j, todo)

			}
			break
		}
	}
}
