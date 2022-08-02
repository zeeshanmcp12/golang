package main

import "fmt"

func main() {
	tasks := []string{"Morning Walk", "Workout", "Learn to code", "Write some lines"}
	// fmt.Println(tasks)
	for j, todo := range tasks {
		fmt.Printf("%v task: %v\n", j, todo)
	}

	var index int

	fmt.Print("Select any number to delete task: ")
	fmt.Scanf("%v", &index)

	// index += 0

	// fmt.Println("Value in index: ", index)

	tasks = append(tasks[:index], tasks[index+1:]...)

	// still working on it
	var updated_task []string = tasks[index:]
	// fmt.Println("Updated tasks: ", updated_task)

	for j, todo := range updated_task {
		fmt.Printf("%v task: %v\n", j, todo)
	}

}
