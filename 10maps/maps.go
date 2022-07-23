package main

import "fmt"

func main() {

	fmt.Println("Revisiting Maps in golang!")

	abbr := make(map[string]string)

	abbr["OOO"] = "Out Of Office"
	abbr["WFH"] = "Work From Home"
	abbr["BRB"] = "Be Right Back"

	// fmt.Println("Abbreviations used in office: ", abbr)

	for key, value := range abbr {
		fmt.Printf("%v -> %v\n", key, value)
	}

}
