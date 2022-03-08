package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Decoding JSON and reading it from file!")

	// ReadFile()
	DecodeJSON()
}

func ReadFile() {
	data, err := ioutil.ReadFile("../challengeExercises/userProfiles_db.json")

	CheckNilErr(err)

	// fmt.Printf("Type of data is %T", data)
	fmt.Println(string(data))

}

func DecodeJSON() {
	fmt.Println("Decoding JSON data!")
	// ReadFile()

	var userData Profile

	file, err := ioutil.ReadFile("../challengeExercises/userProfile.json")
	CheckNilErr(err)

	jsonDataFromFile := []byte(file)

	// fmt.Println(string(jsonDataFromFile))

	// isJSONValid := json.Valid(jsonDataFromFile)

	json.Unmarshal(jsonDataFromFile, &userData)

	fmt.Printf("%#v\n", userData)

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromFile, &myOnlineData)
	// fmt.Printf("%#v\n", myOnlineData)

	for k, val := range myOnlineData {
		fmt.Printf("Key: %v\tValue: %v\tType: %T\n", k, val, val)
	}

}

/*
func DecodeJSON() {
	fmt.Println("Decoding JSON!")

	var userData Profile
	file, err := ioutil.ReadFile("../challengeExercises/userProfiles_db.json")

	CheckNilErr(err)

	jsonDataFromFile := []byte(file)

	json.Unmarshal(jsonDataFromFile, &userData)

	fmt.Printf("%#v\n", userData)

}*/

type Profile struct {
	Name     string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"emailaddress"`
	Age      int    `json:"age"`
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
