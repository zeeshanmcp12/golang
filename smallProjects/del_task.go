package main

import "fmt"

func main() {
	fmt.Println("Handling about deleting task based on index number!")

	tasks := []string{"Morning walk", "Workout", "Learn to Code", "Write some lines", "Help someone"}

	var indexToDel int

	fmt.Println("Select number to delete the task: ")
	for j, val := range tasks {
		fmt.Printf("%v: %v\n", j, val)
	}
	fmt.Scanf("%v", &indexToDel)
	// fmt.Printf("You selected task: %v ", tasks[indexToDel])

	for i := 0 + 1; i < len(tasks); i++ {

		if indexToDel < len(tasks) {
			fmt.Printf("You selected %v: %v ", indexToDel, tasks[indexToDel])

		}
		break

	}

	tasks = append(tasks[:indexToDel], tasks[indexToDel+1:]...)

	fmt.Println("\n-------------------------")

	fmt.Println("Tasks have been updated: ")
	for j, val := range tasks {
		fmt.Printf("%v: %v\n", j, val)
	}
}
