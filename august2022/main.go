package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Building applications in golang!")

	tasks := []string{}

	fmt.Printf("Add tasks: ")

	for i := 0; true; i++ {

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		f_txt := strings.TrimSpace(input)

		if f_txt != "done" {
			if f_txt != "" {
				tasks = append(tasks, f_txt)
				fmt.Printf("Add task or type done: ")
				continue
			} else {
				fmt.Printf("Invalid input, add task or type done: ")
			}
		} else if f_txt == "done" {
			fmt.Println("Thank you for adding tasks: ")
			for j, todo := range tasks {
				fmt.Printf("%v task: %v\n", j, todo)
			}
			break
		} else {
			fmt.Printf("Invalid Input!")
		}

	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Press enter to continue...")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
}
