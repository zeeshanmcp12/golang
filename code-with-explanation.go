package main

import "fmt"

type hotdog int

var x hotdog

var y int

func main() {
	fmt.Println("Code with some explanation")

	x = 42
	fmt.Printf("Type of x -> %T\t Value of x -> %d\n", x, x)

	// Convert and assign CUSTOM type to y
	fmt.Printf("Address of y -> %v\n", &y)
	y := int(x)
	fmt.Printf("Type of y -> %T\t Value of y -> %d\n", y, y)
	fmt.Printf("New Address of y -> %v\n", &y)

	// Difference between y = int(x) and y := int(x)
	// y = int(x) -> convert and assign x to y
	// In this case, the address of y in memory will remain same
	// y := int(x) -> redeclare variable y, convert and assign x to y
	// In this case, the original address of y in memory will be lost
	// Execute above program by changing it

	// Output with y = int(x)
	// Address of y -> 0x10d5b48
	// New Address of y -> 0x10d5b48

	// Output with y := int(x)
	// Address of y -> 0xd65b48
	// New Address of y -> 0xc000012098

}
