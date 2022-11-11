package main

import (
	"encoding/json"
	"fmt"
)

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
	EncodeJson()
}

func EncodeJson() {

	ourKitten := []kitten{
		{"Ruby", 60, "Brownish", "Good Hunter", []string{"Fluffy", "Canned Food"}},
		{"Max", 60, "White", "Keen observer", []string{"Fluffy", "Good food"}},
	}

	// json.Marshal
	// Marshal -> is a method inside json package.
	// Marshal -> will not show json data properly (or in json format) and only show data in single line.
	// data, err := json.Marshal(ourKitten)

	// MarshalIndent -> is like Marshal but applies Indent to format the output. Each JSON element in the output will begin on a new line beginning with prefix followed by one or more copies of indent according to the indentation nesting.
	data, err := json.MarshalIndent(ourKitten, "", "\t")
	// json.MarshalIndent(ourKitten, "", "\t")
	// MarshalIndent -> requires three arguments
	// 		Interface -> struct (which is ourKitten in our case)
	// 		Prefix -> any string which will add in the beginning of new line. In our case it is ""(empty string)
	// 		indent string -> \t -> tab (in our case)

	CheckNilErr(err)

	fmt.Printf("%s\n", data)
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
