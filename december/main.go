package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var name string = "Sakina"

	for i, j := range name {
		fmt.Printf("%v%c\n", i, j)
	}
}
