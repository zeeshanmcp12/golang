package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Slice in golang!")

	tasks := []string{}

	fmt.Print("Add tasks: ")

	for i := 0; true; i++ {

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		f_tasks := strings.TrimSpace(input)

		if f_tasks != "done" {
			if f_tasks != "" {
				tasks = append(tasks, f_tasks)
				fmt.Print("Add task or type done: ")
			} else {
				fmt.Print("Try again, add task or type done: ")
			}
			continue
		} else if f_tasks == "done" {
			fmt.Println("Thank you for adding tasks: ")
			for items, value := range tasks {
				fmt.Printf("No %v: %v\n", items+1, value)
			}
			break
		} else {
			fmt.Println("Invalid input. Try again!")
		}
	}
}
