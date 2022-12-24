package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("User Input and JSON formatting!")

	// Encode data to JSON
	// EncodeJson()

	//Decode data to JSON
	DecodeJson()

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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"username": "Zeeshan",
		"emailaddress": "zee@go.dev",
		"skills": ["Azure","Terraform","Kubernetes"]
	}
	`)

	var userData User

	// one way to decode the json
	isJsonValid := json.Valid(jsonDataFromWeb)

	if isJsonValid {
		fmt.Println("JSON is Valid")
		json.Unmarshal(jsonDataFromWeb, &userData)
		fmt.Printf("%#v\n", userData)
	} else {
		fmt.Println("JSON WAS NOT VALID!")
	}

	// another way to decode the json
	// In some cases we would only want to add data to key value
	var onlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &onlineData)
	fmt.Printf("%#v\n", onlineData)

	for k, val := range onlineData {
		fmt.Printf("Key -> %v, Value -> %v, Type -> %T\n", k, val, val)
	}

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
