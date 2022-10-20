package main

import "fmt"

func main() {

	fmt.Println("Maps in golang!")

	abbr := make(map[string]string)

	abbr["WFH"] = "Work From Home"
	abbr["BRB"] = "Be Right Back"
	abbr["OOO"] = "Out Of Office"

	fmt.Printf("Type of abbr is %T\n", abbr)
	// fmt.Println("Abbreviations are: ", abbr)

	for j, val := range abbr {
		fmt.Printf("Key -> %v, Value -> %v\n", j, val)
	}

}
