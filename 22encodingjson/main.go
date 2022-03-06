package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Encoding Json in golang!")
	EndcodeJSON()
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
		{"Muhammad Zeeshan", 32, "acloudtechie@outlook.com", "abc123+", "acloudtechie.com", []string{"cloud", "azops"}},
		{"Abdullah", 26, "abdullah@outlook.com", "def123+", "acloudtechie.com", []string{"support", "helpdesk"}},
		{"Nasir", 30, "nasir@outlook.com", "ghi123+", "acloudtechie.com", nil},
	}

	// Third package the data in JSON format
	JsonFormat, err := json.MarshalIndent(profileData, "", "\t")
	CheckNilErr(err)

	fmt.Printf("%s", JsonFormat)

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
