package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type User struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Email string   `json:"emailAddress"`
	Skill []string `json:"skills,omitempty"`
}

func main() {
	fmt.Println("Working with Json - Revisit")

	encodeJsonData()
}

func encodeJsonData() {

	userData := []User{
		{"Abdullah", 30, "abd@go.dev", []string{"Keen observer", "Soft spoken"}},
		{"Shani", 31, "sha@go.dev", []string{}},
	}

	encodedJsonData, err := json.MarshalIndent(userData, "", "\t")

	checkNilErr(err)

	fmt.Printf("%s", encodedJsonData)

	createFile(string(encodedJsonData))

}

func createFile(data string) {
	file, err := os.Create("./test2.json")
	checkNilErr(err)

	io.WriteString(file, data)

	defer file.Close()
}

func readFile(filename []byte) {
	fileData, err := ioutil.ReadFile(string(filename))

	checkNilErr(err)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
