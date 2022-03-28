package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Encoding/Decoding Json in golang!")
	// EndcodeJSON()
	DecodeJSON()
}

// To encode data in JSON format:
// First we need to define data structure

type Profile struct {
	FullName string   `json:"name"`
	Age      int      `json:"age"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
	Platform string   `json:"website"`
	Tags     []string `json:"tags,omitempty"`
}

func EndcodeJSON() {
	// Second we need to create the data that has to be package as JSON
	profileData := []Profile{
		{"Muhammad Zeeshan", 32, "acloudtechie@outlook.com", "adfdf", "acloudtechie.com", []string{"cloud", "azops"}},
		{"Abdullah", 26, "abdullah@outlook.com", "adfdf", "acloudtechie.com", []string{"support", "helpdesk"}},
		{"Nasir", 30, "nasir@outlook.com", "asfdf", "acloudtechie.com", nil},
	}

	// Third package the data in JSON format
	JsonFormat, err := json.MarshalIndent(profileData, "", "\t")
	CheckNilErr(err)

	fmt.Printf("%s", JsonFormat)

}

func DecodeJSON() {
	fmt.Println("Consuming JSON data!")

	jsonStruct := []byte(`
	{
		"name": "Muhammad Zeeshan",
		"age": 32,
		"email": "acloudtechie@outlook.com",
		"website": "acloudtechie.com",      
		"tags": ["cloud","azops"]
	}
	`)

	var jsonData Profile

	isJSONValid := json.Valid(jsonStruct)

	if isJSONValid {

		fmt.Println("JSON is Valid!")
		json.Unmarshal(jsonStruct, &jsonData) // here as 1st parameter, we need to pass "slice of type byte" which is jsonStruct in our case.
		fmt.Printf("%#v\t", jsonData)

	} else {
		fmt.Println("JSON is not valid!")
	}

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
