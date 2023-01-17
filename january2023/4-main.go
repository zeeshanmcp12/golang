package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Todo App!")

	tasks := []string{}

	for true {

		fmt.Printf("Enter Tasks: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		f_txt := strings.TrimSpace(input)

		if f_txt != "done" {
			if f_txt != "" {
				tasks = append(tasks, f_txt)
				fmt.Printf("Enter task or type done: \n")
			} else {
				fmt.Printf("Invalid input! Enter task or type done: \n")
			}
			continue
		} else if f_txt == "done" {
			for i, val := range tasks {
				fmt.Printf("%v -> %v\n", i, val)
			}
		} else {
			fmt.Printf("Invalid input! Enter task or type done: \n")
		}
		break

	}
}
