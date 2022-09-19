package main

import "fmt"

func main() {
	fmt.Println("Revisiting slice in golang!")

	tasks := []string{"Offer prayer fajr", "Workout", "Learn to code", "Help someone"}
	fmt.Printf("Type of tasks is %T", tasks)
	fmt.Println()

	for key, value := range tasks {
		fmt.Printf("Task#%v: %v\n", key, value)
	}
}
