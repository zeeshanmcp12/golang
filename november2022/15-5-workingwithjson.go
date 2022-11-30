package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string   `json:"name"`
	Age    int      `json:"age"`
	Email  string   `json:"email"`
	Skills []string `json:"skills,omitempty"`
}

func main() {
	fmt.Println("Working with json")

	encodeJson()
}

func encodeJson() {

	userData := []User{
		{"Abdullah", 29, "abd@go.dev", []string{"Scrum", "Agile Process"}},
		{"Asim", 26, "asm@go.dev", []string{}},
	}

	encodedData, err := json.MarshalIndent(userData, "", "\t")

	checkNilErr(err)

	fmt.Printf("%s\n", encodedData)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
