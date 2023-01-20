package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Todo app logic writing in NADRA office :)")

	tasks := []string{}

	fmt.Printf("Add Task: ")

	for true {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		f_txt := strings.TrimSpace(input)

		if f_txt != "done" {
			if f_txt != "" {
				tasks = append(tasks, f_txt)
				fmt.Printf("Add task or type done: ")
			} else {
				fmt.Printf("Invalid input! Add task or type done: ")
			}
			continue
		} else if f_txt == "done" {
			fmt.Printf("Your Tasks for today: \n")
			for i, val := range tasks {
				fmt.Printf("%v -> %v\n", i+1, val)
			}
			break
		}
	}
}
