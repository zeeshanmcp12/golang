package main

import "fmt"

func main() {
	fmt.Println("Remove slice based on index number!")

	var todos = []string{"Morning Walk", "Workout", "Learn to code", "Write something", "Help someone"}
	fmt.Println("Length of slice: ", len(todos))

	for j, todo := range todos {
		fmt.Printf("%v task: %v\n", j, todo)
	}

	var index int
	fmt.Print("Select any number to delete the task: ")
	fmt.Scanf("%v", &index)

	todos = append(todos[:index], todos[index+1:]...)

	fmt.Println("--------------------")
	fmt.Println("Todos updated:")
	for j, todo := range todos {
		fmt.Printf("%v task: %v\n", j, todo)
	}

}
