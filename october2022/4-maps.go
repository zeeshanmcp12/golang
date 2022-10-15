package main

import "fmt"

func main() {
	fmt.Println("Find abbreviation!")

	abbr := make(map[string]string)

	abbr["OOO"] = "Out of Office"
	abbr["WFH"] = "Work From Home"
	abbr["BRB"] = "Be Right Back"

	// fmt.Println("OOO stands for ->", abbr["OOO"])

	// for j, item := range abbr {
	// 	fmt.Printf("%v stands for %v\n", j, item)
	// }

	var abbrToSel string

	fmt.Printf("Find abbreviations from below: ")
	fmt.Println()

	for j := range abbr {
		fmt.Printf("[%v]: ", j)
	}

	fmt.Scanf("%v", &abbrToSel)

	_, ok := abbr[abbrToSel]

	if ok == true {
		fmt.Println("Abbreviation found!")
		// for j := range abbr {
		// }
		fmt.Printf("%v stands for %v", abbrToSel, abbr[abbrToSel])

	} else {
		fmt.Println("Abbreviation not found!")
	}
}
