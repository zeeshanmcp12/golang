package main

import "fmt"

/*
Using the code from the previous exercise,
1. At the package level scope, assign the following values to the three variables
a. for x assign 42
b. for y assign “James Bond”
c. for z assign true
2. in func main
a. use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the
returned value of TYPE string using the short declaration operator to a
VARIABLE with the IDENTIFIER “s”
b. print out the value stored by variable “s”

*/

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Println("----- Ninja Level 1 ----- Coding Exercise #3 -----")

	s := fmt.Sprintf("x -> %d, y -> %s, z -> %t", x, y, z)

	fmt.Printf("Type of s variable is %T\n", s)

	fmt.Println(s)

}
