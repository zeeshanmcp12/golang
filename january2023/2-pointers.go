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

	p := &i
	fmt.Println("New variable p -> (var p = &i)")

	fmt.Printf("Value of p -> %v is the address of i and value of *p is the value at that address -> %v\n", p, *p)
	fmt.Println("--------------------------------------")
	fmt.Println("Let's change the value of i by assigning new value to *p -> *p = 21. It should change the value of i to 21")

	*p = 21
	fmt.Printf("Value of *p -> %v\n", *p)
	fmt.Printf("Value of i -> %v, hence the value of i has changed \n", i)

	fmt.Println("--------------------------------------")
	fmt.Println("--------------------------------------")
	p = &j
	fmt.Println("New variable (var p = &j)")
	fmt.Println("Let's change the value of j by assigning it to p (var)!")

	*p = *p / 73
	fmt.Println("*p = *p / 73")
	fmt.Println("*p = *p (value at j -> 2701) / 73")
	fmt.Printf("Value of j -> %v, hence the value of j has changed \n", j)

	fmt.Println("--------------------------------------")
	fmt.Println(" ----- Now after performing changes -----")
	fmt.Printf("Value of i -> %v and j -> %v.\n", i, j)

	fmt.Println("--------------------------------------")
	fmt.Println(" ---------- Why do we need pointers ???---------- ")
	fmt.Printf("When we can write j = j / 73\ninstead of *p = *p / 73...\n")
	fmt.Printf("It's important to store variable at one place and use here and there (or wherever we want to use and modify)\n")
	fmt.Printf("We can share the variable and update it in multiple places instead of copy the variable everytime we need to use or update it.")

}
