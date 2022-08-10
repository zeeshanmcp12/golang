package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang!")

	abbr := make(map[string]string)

	abbr["OOO"] = "Out Of Office"
	abbr["WFH"] = "Work From Home"
	abbr["BRB"] = "Be Right Back"

	// fmt.Println(abbr)

	// for j, list := range abbr {
	// 	fmt.Printf("%v -> %v\n", j, list)
	// }

	// fmt.Printf("Type of abbr map is: %T", abbr)

	// abbr = delete(abbr["WFH"])

	var stringToSel string

	fmt.Println("Select abbreviation to find out the word: ")
	for j := range abbr {
		// fmt.Printf("null %v\n", list)
		fmt.Printf("%v ", j)

	}
	fmt.Println()
	fmt.Scanf("%v", &stringToSel)
	fmt.Printf("%v stands for %q", stringToSel, abbr[stringToSel])

}
