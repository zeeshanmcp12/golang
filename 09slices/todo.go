package main

import "fmt"

func main() {
	fmt.Println("Writing logic for add/delete task!")

	fmt.Print("Select any number: [1]Add Task, [2]Delete Task: ")

	var menuIndex int

	fmt.Scanf("%v", &menuIndex)

	if menuIndex == 1 {
		fmt.Println("Please add task!")
		// continue

	} else if menuIndex == 2 {
		fmt.Println("Please delete task!")
		// break
	}
}
