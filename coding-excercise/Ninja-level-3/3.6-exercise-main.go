package main

import "fmt"

// Create a program that shows the “if statement” in action.

func main() {
	fmt.Println("Exercise 6 - Ninja level 3")

	name := "Abdullah"

	bs := []byte(name)

	fmt.Printf("%s -> %d \n", bs, bs)

	for i := 0; i < len(bs); i++ {
		fmt.Printf("%s -> %d \n", string(bs[i]), bs[i])

	}

	if name == "Abdullah" {
		fmt.Println("String found -> ", name)
	}
}
