package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("User Input and JSON formatting!")

	EncodeJson()

}

type User struct {
	Name   string   `json:"username"`
	Email  string   `json:"emailaddress"`
	Skills []string `json:"skills,omitempty"`
}

func EncodeJson() {

	userData := []User{
		{"Zeeshan", "zee@go.dev", []string{"Azure", "Terraform", "Kubernetes"}},
		{"Abdullah", "abd@go.dev", []string{"AWS", "Bicep", "Kubernetes"}},
	}

	jsonData, err := json.MarshalIndent(userData, "", "\t")
	checkNilErr(err)

	createFile(jsonData)

}

func createFile(data []byte) {
	file, err := os.Create("./user.json")

	checkNilErr(err)

	io.WriteString(file, string(data))

	defer file.Close()

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
