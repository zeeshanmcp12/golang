package main

import "fmt"

func main() {
	fmt.Println("Working with Pointers in golang!")
	fmt.Println()
	fmt.Println("--------------------------------------")

	var i, j = 42, 2701

	fmt.Printf("Value of i -> %v and j -> %v.\n", i, j)
	fmt.Printf("Address of i -> %v and j -> %v in memory.\n", &i, &j)

	fmt.Println("--------------------------------------")
	fmt.Println("--------------------------------------")

	var p = &i
	fmt.Println("New variable p -> (var p = &i)")

	fmt.Printf("Value of p -> %v is the address of i and value of *p is the value at that address -> %v", p, *p)

}
