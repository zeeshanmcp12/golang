package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Slices in golang!")

	var todos = []string{}
	// fmt.Printf("Type of todos is %T\n", todos)

	// todos = append(todos, "item1", "item2", "item3")
	// fmt.Println("Todos are ", todos)

	// var todoItem string

	fmt.Print("Enter todo item: ")

	for i := 0; true; i++ {
		// fmt.Scanf("%s\n", &todoItem)

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		f_text := strings.TrimSpace(input)

		if f_text != "done" {
			if f_text != "" {
				todos = append(todos, f_text)
				fmt.Println("Type 'done' or add another task: ")
				continue
			} else {
				fmt.Println("Invalid input, type 'done' or add another task: ")
			}
		} else if f_text == "done" {
			// fmt.Println("Todos updated: ", todos)
			for j, todo := range todos {
				fmt.Printf("%v todo: %s\n", j+1, todo)
			}
			break

		} else {
			fmt.Print("Invalid input!")
		}
	}
}
