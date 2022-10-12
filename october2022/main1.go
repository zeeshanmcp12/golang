package main

import "fmt"

func main() {
	fmt.Println("Resuming again!")
	fmt.Println("Working with slice in golang!")

	someItems := []string{"Item One", "Item Two", "Item Three", "Item Four"}

	for key, item := range someItems {
		fmt.Printf("Key : %v, Item : %v\n", key, item)
	}
}
