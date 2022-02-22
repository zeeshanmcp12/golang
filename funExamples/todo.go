package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("------Todo app------")
	todos := []string{}
	fmt.Printf("Add todos: ")

	for i := 0; true; i++ {

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		f_text := strings.TrimSpace(input) // f_text refers to 'final text' after triming space

		if f_text != "Done" {
			if f_text != "" {
				fmt.Printf("Enter task or type 'Done': ")
				todos = append(todos, f_text)
				continue
			} else {
				fmt.Printf("Invalid Text! Enter task or type 'Done': ")
			}

		} else if f_text == "Done" {

			fmt.Println()
			// j, todo -> to get the index and it's value alongside
			for j, todo := range todos {
				fmt.Printf("Task %v %v\n", j+1, todo)
			}
			fmt.Println()
			break

		}
	}
}
