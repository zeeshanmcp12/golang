package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Revisiting concepts by writing todo app!")

	tasks := []string{}

	fmt.Printf("Enter tasks: ")

	for i := 0; true; i++ {

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		f_text := strings.TrimSpace(input)

		if f_text != "" {
			if f_text != "done" {
				tasks = append(tasks, f_text)
				fmt.Printf("Enter task or type done: ")

			} else {
				fmt.Printf("Enter task or type done! ")

			}
			continue
		} else if f_text == "done" {
			for key, item := range tasks {
				fmt.Printf("Key: %v, Item: %v\n", key, item)
			}
			break
		} else {
			fmt.Printf("Invalid input, please try again!")
		}

	}

}
