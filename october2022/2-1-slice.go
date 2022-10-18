package main

import "fmt"

func main() {
	fmt.Println("Remove slice based on index")

	var indexSelection int

	tasks := []string{"Play Cricket", "Play Foosball", "Play Table Tennis", "Play Games"}

	for j, v := range tasks {
		fmt.Printf("Index -> %v, Value -> %v\n", j, v)
	}

	fmt.Println("--------------------")
	fmt.Print("Select any number from above list to remove the task (0-3): ")

	fmt.Scanf("%v", &indexSelection)

	fmt.Printf("You selected %q\n", tasks[indexSelection])

	tasks = append(tasks[:indexSelection], tasks[indexSelection+1:]...)

	fmt.Println("Updated tasks: ")
	fmt.Println("--------------------")

	for j, v := range tasks {
		fmt.Printf("Index -> %v, Value -> %v\n", j, v)
	}

}
