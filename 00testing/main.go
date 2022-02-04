package main

import "fmt"

func main() {
	fmt.Println("New file for practice")
	var todos = []string{"ReciteQuran", "Workout", "LearnToCode"}
	fmt.Println("Todos are:", todos)
	fmt.Println("Items in todos:", len(todos))
	// fmt.Println()
	fmt.Println("------------------------------")
	fmt.Println("Add items in array")
	fmt.Println("------------------------------")
	todos = append(todos, "SprintTask", "WriteSomeWords")
	fmt.Println("Updated todos:", todos)
	fmt.Println("----------Start Slicing-----------")
	// todos = append(todos[1:])
	// fmt.Println("After slicing: ", todos)
	todos = append(todos[:4])
	fmt.Println("from 0th to 4(non-inclusive)", todos)

}
