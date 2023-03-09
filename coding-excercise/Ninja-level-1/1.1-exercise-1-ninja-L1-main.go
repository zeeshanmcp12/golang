package main

import "fmt"

/*
Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the
IDENTIFIERS “x” and “y” and “z”
a. 42
b. “James Bond”
c. true
2. Now print the values stored in those variables using
a. a single print statement
b. multiple print statements
*/

func main() {
	fmt.Println("----- Ninja Level 1 ----- Coding Exercise #1 -----")

	age := 42
	username := "James Bond"
	isHeHero := true

	fmt.Println("Yes that's", isHeHero, "His name is", username, "and he is", age, "years old.")
	fmt.Println(username)
	fmt.Println(age)
	fmt.Println(isHeHero)
}
