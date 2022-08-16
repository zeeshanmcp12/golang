package main

import (
	"fmt"
)

func main() {
	fmt.Println("Find abbreviations!")

	abbr := make(map[string]string)

	abbr["OOO"] = "Out Of Office"
	abbr["WFH"] = "Work From Home"
	abbr["BRB"] = "Be Right Back"

	// fmt.Println(abbr)

	// for j, val := range abbr {
	// 	fmt.Printf("%v stands for %q\n", j, val)
	// }

	var abbrToSelect string
	fmt.Printf("Find words for these abbreviations: ")
	for j := range abbr {
		fmt.Printf("%q ", j)
	}
	// fmt.Println()
	fmt.Scanf("%v", &abbrToSelect)

	_, ok := abbr[abbrToSelect]

	fmt.Println("abbreviation is available so", ok)

	if ok == true {
		fmt.Printf("%v stands for %v", abbrToSelect, abbr[abbrToSelect])
	} else {
		fmt.Println("Abbreviation doesn't match with data!")
	}
}
