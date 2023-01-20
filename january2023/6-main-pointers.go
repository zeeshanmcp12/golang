package main

import "fmt"

func main() {
	fmt.Println("Working with pointers in golang!")

	var x, y = 20, 1010
	fmt.Println(x, y)

	fmt.Printf("Address of x -> %v\nAddress of y -> %v\n", &x, &y)

	a := &x
	fmt.Printf("Address of a -> %v\nValue at that address that the pointer is pointing to -> %v\n", a, *a)

	fmt.Println("Let's change the value of x through a and not directly x.")

	*a = 30
	fmt.Printf("Value of x has changed to %v", x)
}
