package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string
	Age   int
	Email string
	Skill []string
}

func main() {
	fmt.Println("Working with Json - Revisit")

	encodeJsonData()
}

func encodeJsonData() {

	userData := []User{
		{"Abdullah", 30, "abd@go.dev", []string{"Keen observer", "Soft spoken"}},
		{"Shani", 31, "sha@go.dev", []string{"Good learner", "Soft spoken"}},
	}

	encodedJsonData, err := json.MarshalIndent(userData, "", "\t")

	checkNilErr(err)

	fmt.Printf("%s", encodedJsonData)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
