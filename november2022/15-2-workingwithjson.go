package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type User struct {
	Name     string   `json:"username"`
	Age      int      `json:"age"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
	Skills   []string `json:"skills,omitempty"` // This omitempty will remove the empty value from output
}

func main() {
	fmt.Println("Working with json in golang!")

	encodeJsonData()
}

func encodeJsonData() {

	userData := []User{
		{"Abdullah", 25, "abd@go.dev", "abc123+", []string{"Java", "AWS"}},
		{"Waqas", 27, "waq@go.dev", "abc1234+", []string{".NET", "Rest API"}},
		{"Zulfi", 27, "zul@go.dev", "abc@123+", []string{}},
	}

	encodedData, err := json.MarshalIndent(userData, "", "\t")

	CheckNilErr(err)

	fmt.Printf("%s\n", encodedData)

	// Converting json encoded data to string so we can use it in createFile function
	// Because io.WriteString only accepts string as an argument
	convertToString := string(encodedData)

	file := "./test.json"

	// Calling createFile function
	createFile(file, convertToString)

}

// Generate json file
func createFile(filename string, jsondata string) {
	file, err := os.Create(filename)
	CheckNilErr(err)

	io.WriteString(file, jsondata)

	defer file.Close()
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
