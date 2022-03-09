package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Logic not working
func main() {
	fmt.Println("Reading JSON data from file!")

	EncodeJSON()
}

func EncodeJSON() {

	ipRange := ReadFile()

	jsonData, _ := json.MarshalIndent(ipRange, "", "\t")

	// fmt.Println(jsonData)
	// fmt.Printf("%s", string(jsonData))
	fmt.Println(string(jsonData))

}

func ReadFile() []byte {

	file, err := ioutil.ReadFile("../challengeExercises/userProfile.json")
	CheckNilErr(err)

	JSONData := []byte(file)

	return JSONData
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
