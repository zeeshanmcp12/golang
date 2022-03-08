package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Decoding JSON and reading it from file!")

	ReadFile()
}

func ReadFile() {
	data, err := ioutil.ReadFile("../challengeExercises/userProfile.json")

	CheckNilErr(err)

	// fmt.Printf("Type of data is %T", data)
	fmt.Println(string(data))

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
