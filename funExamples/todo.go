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

	for i := 0; true; i++ {

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Add todo: ")
		input, _ := reader.ReadString('\n')
		// fmt.Println(input)
		f_text := strings.TrimSpace(input)
		todos = append(todos, f_text)
		fmt.Println(f_text)

		if f_text != "Done" {
			fmt.Printf("Enter another task or type 'Done' when you're done: ")
			moreTodos, _ := reader.ReadString('\n')
			mt_f_text := strings.TrimSpace(moreTodos)
			todos = append(todos, mt_f_text)
			fmt.Printf("Current tasks %v\n", todos)
			continue

		} else if f_text == "Done" {
			fmt.Printf("Your Todos %v\n", todos)
			break
		} else {
			fmt.Println("Invalid input!")
		}

	}

	// fmt.Println(todos)

}
