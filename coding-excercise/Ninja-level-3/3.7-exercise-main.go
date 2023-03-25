package main

import "fmt"

// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

func main() {
	fmt.Println("Exercise 7 - Ninja level 3")

	name := "Abdullah"

	bs := []byte(name)

	fmt.Printf("%s -> %d \n", bs, bs)

	for i := 0; i < len(bs); i++ {
		fmt.Printf("%s -> %d \n", string(bs[i]), bs[i])

	}

	if name == "Abdullah" {
		fmt.Println("String found -> ", name)
	} else if name != "Abdullah" {
		fmt.Println("String not found", name)
	} else {
		fmt.Println("Invalid input")
	}
}
