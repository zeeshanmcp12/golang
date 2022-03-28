package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("Working with JSON. Encoding/Decoding")
	EncodingJson()
}

func EncodingJson() {

	userProfile := []User{
		{"Muhammad Zeeshan", 32, "+", "zeeshan@dev.io"},
		{"Hussain Raza", 28, "+", "hussain@dev.io"},
		{"Abdullah Shafiq", 25, "+", "raza@dev.io"},
	}

	finalJson, err := json.MarshalIndent(userProfile, "", "\t")
	CheckNilErr(err)

	fmt.Printf("%s\n", finalJson)

}

type User struct {
	Name        string `json:"fullname"`
	Age         int    `json:"age"`
	Password    string `json:"-"`
	ContactInfo string `json:"email"`
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
