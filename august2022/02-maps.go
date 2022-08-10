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

	abbr["SL"] = "Sick Leave"

	fmt.Println("List updated!")

	for j, list := range abbr {
		fmt.Printf("%v -> %v\n", j, list)
	}

}
