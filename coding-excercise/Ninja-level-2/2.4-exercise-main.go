package main

import "fmt"

/*
Write a program that
● assigns an int to a variable
● prints that int in decimal, binary, and hex
● shifts the bits of that int over 1 position to the left, and assigns that to a variable
● prints that variable in decimal, binary, and hex

*/

var intNumber int

func main() {
	fmt.Println("Exercise 4 in Ninja level 2")

	intNumber = 42

	fmt.Printf("Binary -> %b\t Decimal -> %d\t Hex -> %#x\n", intNumber, intNumber, intNumber)

	intNumberLeftShifted := intNumber << 1
	fmt.Printf("Binary -> %b\t Decimal -> %d\t Hex -> %#x", intNumberLeftShifted, intNumberLeftShifted, intNumberLeftShifted)

}

/*
 101010
1010100
64+16+4=84
*/
