package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Combining all concepts")

	tasks := []string{}
	fmt.Print("Add task: ")

	for i := 0; true; i++ {

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		f_txt := strings.TrimSpace(input)

		if f_txt != "done" {
			if f_txt != "" {
				tasks = append(tasks, f_txt)
				fmt.Print("Type 'done' or add task: ")
			} else {
				fmt.Print("No input, type 'done' or add task: ")
			}
			continue
		} else if f_txt == "done" {
			fmt.Println("Thank you for adding tasks: ")
			for j, val := range tasks {
				fmt.Printf("Task# %v -> %v \n", j+1, val)
			}
			break
		} else {
			fmt.Println("Invalid input!")
		}

	}
}
