package main

import "fmt"

func main() {
	tasks := []string{"1: Task1", "2: Task2", "3: Task3", "4: Task4"}
	fmt.Println(tasks)

	var index int

	fmt.Print("Select any number to delete task: ")
	fmt.Scanf("%v", &index)

	tasks = append(tasks[:index], tasks[index+1:]...)
	for j, todo := range tasks {
		fmt.Printf("%v task: %v\n", j, todo)
	}

}
