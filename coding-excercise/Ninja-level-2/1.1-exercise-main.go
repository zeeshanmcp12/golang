package main

import "fmt"

// Write a program that prints a number in decimal, binary, and hex

// var number int = 23

func main() {

	number := 23

	fmt.Println("Ninja Level 2 Exercise")
	// fmt.Printf("Number -> %f in decimal format.", number)
	fmt.Printf("Number -> %b in binary format.\n", number)
	fmt.Printf("Number -> %d in decimal format.\n", number)
	fmt.Printf("Number -> %#x in hexadecimal format.\n", number)

	fmt.Printf("%d\t%b\t%#x\n", number, number, number)
}
