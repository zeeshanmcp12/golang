package main

import "fmt"

func main() {
	fmt.Println("Using `raw string literal`")

	multilineString := `"Multi line
	string"`
	fmt.Println(multilineString)

	fmt.Println("Convert string to byte format")

	myName := "Abdullah"

	byteString := []byte(myName)
	fmt.Printf("Type of byteString var -> %T\n", byteString)
	fmt.Println(byteString)

	for _, val := range byteString {
		// fmt.Println(i, val)
		fmt.Printf(" -> %d -> %b\n", val, val)
	}
}

/*
String to slice of byte to binary
Abdullah -> string
[65 98 100 117 108 108 97 104] -> slice of byte of string "Abdullah"

A -> 65 -> 1000001
b -> 98 -> 1100010
d -> 100 -> 1100100
u -> 117 -> 1110101
l -> 108 -> 1101100
l -> 108 -> 1101100
a -> 97 -> 1100001
h -> 104 -> 1101000
*/
