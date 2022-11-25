package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type User struct {
	Name   string
	Age    int
	Email  string
	Skills []string
}

func main() {
	fmt.Println("Working with Json and save in file!")

	encodeJsonData()
}

func encodeJsonData() {

	userData := []User{
		{"Abdullah", 30, "abd@go.dev", []string{"terraform", "azure"}},
		{"Ahmed", 25, "ahm@go.dev", []string{"bicep", "aws", "python"}},
	}

	encodedData, err := json.MarshalIndent(userData, "", "\t")
	checkNilErr(err)

	fmt.Printf("%s", encodedData)

	writeToFile(string(encodedData))

}

func writeToFile(jsonData string) {

	file, err := os.Create("./test-json-again.json")

	checkNilErr(err)

	io.WriteString(file, jsonData)

	defer file.Close()

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
