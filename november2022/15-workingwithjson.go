package main

import "fmt"

// While working with JSON, these are the points
// Create structure
// Encode data into json
// 	To encode data into json we need a function
// 	Inside this function, we need to create slice and use struct as a slice

type kitten struct {
	Name  string
	Age   int
	Color string
	Skill string
	Food  []string
}

func main() {
	fmt.Println("Working with JSON!")
}

func EncodeJson() {

	ourKitten := []kitten{
		{"Ruby", 60, "Brownish", "Good Hunter", []string{"Fluffy", "Canned Food"}},
		{"Max", 60, "White", "Keen observer", []string{"Fluffy", "Good food"}},
	}
}
