package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Handling tasks using slice in golang!")

	tasks := []string{}

	fmt.Println("Add tasks:")

	for i := 0; true; i++ {

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		task_t := strings.TrimSpace(input)

		if task_t != "done" {
			if task_t != "" {
				tasks = append(tasks, task_t)
				fmt.Println("Type 'done' or add task: ")
				continue
			} else {
				fmt.Println("Invalid input, type 'done' or add task: ")
			}
		} else if task_t == "done" {
			fmt.Println("Thank you for adding tasks:")
			for j, val := range tasks {
				fmt.Printf("%v: %v\n", j, val)
			}
			break
		} else {
			fmt.Println("Invalid input!")
		}
	}
}
