package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Working with JSON!")

	// Fifth, we need to call the EncodeJson function here
	EncodeJson()
}

// First we need to create structure
type User struct {
	Name  string
	Age   int
	Email string
	Skill []string
}

// Second we need to write function to encode JSON
func EncodeJson() {

	// Third we need to use struct of type User which we will convert to json
	userData := []User{
		{"Zeeshan", 31, "zee@go.dev", []string{"Azure", "Terraform", "Kubernetes"}},
		{"Abdullah", 30, "abd@go.dev", []string{"AWS", "Terraform", "Jenkins"}},
		{"Noor", 25, "nor@go.dev", nil},
	}

	// Forth we need to encode the data into JSON
	// json -> is a library
	// MarshalIndent is a method in json library
	// 	MarshalIndent is being used to format the JSON data
	jsonData, err := json.MarshalIndent(userData, "", "\t")
	checkNilErr(err)

	fmt.Printf("%s\n", jsonData)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
