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
	fmt.Printf("Add tasks: ")

	for true {

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		f_text := strings.TrimSpace(input)

		if f_text != "done" {
			if f_text != "" {
				tasks = append(tasks, f_text)
				fmt.Printf("Enter task or type done: ")
			} else {
				fmt.Printf("Text cannot be empty, enter task or type done: ")
			}
			continue
		} else if f_text == "done" {
			fmt.Println("Here are your tasks: ")
			for i, val := range tasks {
				fmt.Printf("Task %v: %v\n", i+1, val)
			}
		} else {
			fmt.Println("Invalid Input. Try Again!")
		}
		break
	}
}

func checkNilErr(err error) {
	panic(err)
}
