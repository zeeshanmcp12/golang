package main

import "fmt"

func main() {
	fmt.Println("Working with Pointers in golang!")
	fmt.Println("--------------------------------------")

	var i, j = 42, 2701

	fmt.Printf("Value of i -> %v and j -> %v.\n", i, j)
	fmt.Printf("Address of i -> %v and j -> %v in memory.\n", &i, &j)

	fmt.Println("--------------------------------------")
	fmt.Println("--------------------------------------")

	var p = &i
	fmt.Println("New variable p -> (var p = &i)")

	fmt.Printf("Value of p -> %v is the address of i and value of *p is the value at that address -> %v\n", p, *p)
	fmt.Println("--------------------------------------")
	fmt.Println("Let's change the value of i by assigning new value to *p -> *p = 21. It should change the value of i to 21")

	*p = 21
	fmt.Printf("Value of *p -> %v\n", *p)
	fmt.Printf("Value of i -> %v, hence the value of i has changed \n", i)

	fmt.Println("--------------------------------------")
	fmt.Println("--------------------------------------")
	fmt.Println("Let's change the value of ")

}
