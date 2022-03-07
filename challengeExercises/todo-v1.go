package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("----- Todo App V1 -----")

	appItems := []string{"'add task'", "'show task'", "'delete task'"}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Select any 1 from these options, %v: ", appItems)

	input, _ := reader.ReadString('\n')
	f_text := strings.TrimSpace(input)

	switch f_text {
	case "add task":
		fmt.Println("Add new task")
	case "show task":
		fmt.Println("Show your tasks")
	case "delete task":
		fmt.Println("Mark 'done' as complete")
	default:
		fmt.Println("Enter correct value!")

	}

}
