package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name   string
	Age    int
	Email  string
	Skills []string
}

func main() {
	fmt.Println("Working with json")

	encodeJson()
}

func encodeJson() {

	userData := []User{
		{"Abdullah", 29, "abd@go.dev", []string{"Scrum", "Agile Process"}},
		{"Asim", 26, "asm@go.dev", []string{"Dev"}},
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
