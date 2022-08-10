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

	for j, list := range abbr {
		fmt.Printf("%v -> %v\n", j, list)
	}

	// fmt.Printf("Type of abbr map is: %T", abbr)

	delete(abbr, "WFH")

	fmt.Println("key deleted!")
	// fmt.Println(abbr)
	for j, list := range abbr {
		fmt.Printf("%v -> %v\n", j, list)
	}

}
