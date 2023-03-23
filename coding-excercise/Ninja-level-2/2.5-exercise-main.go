package main

import "fmt"

/*
Create a variable of type string using a raw string literal. Print it
*/

func main() {
	fmt.Println("Ninja level 2 - Exercise 5")

	rawLiteralString := `Hello from line one 
	World in line two
	we are at line 3
	this is possible because of
	"raw string literal"`

	fmt.Println(rawLiteralString)
}
