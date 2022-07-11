package main

import (
	"fmt"
)

func main() {
	fmt.Println("testing scanf function")

	var text string

	fmt.Print("Enter any text: ")
	// fmt.Scanf("%s\n", &text)

	count, err := fmt.Scanf("%q\n", &text)

	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Given input: ", text)
	fmt.Println("Count: ", count)
	fmt.Println("Error: ", err)

}
