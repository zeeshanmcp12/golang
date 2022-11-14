package main

import (
	"encoding/json"
	"fmt"
)

type Kitten struct {
	Name  string   `json:"petName"`
	Age   int      `json:"age"`
	Color string   `json:"color"`
	Food  []string `json:"petFood"`
	Skill []string `json:"skills,omitempty"`
	// Skill string
}

func main() {
	fmt.Println("Working with Json and it's formatting")
	EncodeJson()
}

func EncodeJson() {

	myKitten := []Kitten{
		{"Ruby", 2, "Brownish", []string{"Fluffy", "Treat"}, nil},
		{"Max", 2, "White", []string{"Fluffy, Nothing New"}, []string{"hunting", "keen observer"}},
	}

	encodedData, err := json.MarshalIndent(myKitten, "", "\t")
	CheckNilErr(err)

	fmt.Printf("%s\n", encodedData)

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
