package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Todo app")
	todos := []string{}
	fmt.Printf("Add todo: ")

	for i := 0; true; i++ {

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		f_text := strings.TrimSpace(input)

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
			fmt.Printf("Your Todos %v\n", todos)
			fmt.Println()
			break

		}
	}
}
