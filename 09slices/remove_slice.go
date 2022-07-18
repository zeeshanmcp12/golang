package main

import "fmt"

func main() {
	fmt.Println("Remove slice item based on index number!")

	var todos = []string{"Offer pray fajr", "Morning walk", "Workout", "Recite Quran", "Learn to code"}
	for j, todo := range todos {
		fmt.Printf("%v task: %v\n", j, todo)
	}

	var index int
	fmt.Print("Select number to delete the task: ")
	fmt.Scanf("%v", &index)

	todos = append(todos[:index], todos[index+1:]...)
	fmt.Println(todos)
	for j, todo := range todos {
		fmt.Printf("%v task: %v\n", j, todo)
	}
}
