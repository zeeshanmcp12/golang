package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Todo App Logic!")

	tasks := []string{}

	for true {

		fmt.Print("Add Task: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		ftxt := strings.TrimSpace(input)

		if ftxt != "done" {
			if ftxt != "" {
				tasks = append(tasks, ftxt)
				fmt.Printf("Type done or Enter task: \n")

			} else {
				fmt.Printf("Invalid input, Type done or Enter task: \n")
			}
			continue
		} else if ftxt == "done" {
			fmt.Printf("Thank you for adding tasks: \n")
			for i, val := range tasks {
				fmt.Println(i, "->", val)
			}
		} else {
			fmt.Println("Invalid Input, Try Again!")
		}
		break
	}

}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
