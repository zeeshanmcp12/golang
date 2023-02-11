package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Revisiting previous concepts!")

	fmt.Printf("Add task: ")
	tasks := []string{}

	for true {

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		ftxt := strings.TrimSpace(input)

		if ftxt != "done" {
			if ftxt != "" {
				tasks = append(tasks, ftxt)
				fmt.Printf("Type done or add task: ")
			} else {
				fmt.Printf("Invalid input, type done or add task: ")
			}
			continue
		} else if ftxt == "done" {
			fmt.Println("Thank you for adding tasks: ")
			for i, val := range tasks {
				fmt.Printf("%v -> %v\n", i+1, val)
			}
		} else {
			fmt.Println("Invalid input. Try Again!")
		}
		break

	}

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
