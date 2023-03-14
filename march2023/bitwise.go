package main

import "fmt"

func main() {
	fmt.Println("Bitwise Operators")
	var num1, num2 = 16, 22

	// bitwise AND
	// result := num1 & num2

	// bitwise OR
	// result := num1 | num2

	// bitwise XOR
	// bitwiseXOR := num1 ^ num2

	// bitwise left-shift
	bitwiseLeftShift := num1 << num2

	fmt.Printf("%b\n%b\n", num1, num2)
	// fmt.Println(result)
	// fmt.Println(bitwiseXOR)
	fmt.Println(bitwiseLeftShift)

}

/*

0011 &
0101
0001 -> 1

01100
10110
00100 -> 4

10000
10110
10000 -> 16
*/

/*
10000
10110
10110 -> 22


*/
