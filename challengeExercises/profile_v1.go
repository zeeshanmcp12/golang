package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Working with files, struct and json encoding!")
	fmt.Printf("Enter Username: ")
	username := strings.TrimSpace(onlyInput())

	fmt.Printf("Enter Password: ")
	password := strings.TrimSpace(onlyInput())

	fmt.Printf("Enter Email Address: ")
	email := strings.TrimSpace(onlyInput())

	fmt.Printf("Enter your Age: ")
	agestoi := userInput()

	userProile := Profile{username, password, email, agestoi}

	// fmt.Printf("Type of struct is %T\n", userProile)

	// structIterator(userProile)

	value := reflect.ValueOf(userProile)

	typeOfStruct := value.Type()

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfStruct.Field(i).Name, value.Field(i).Interface())

	}

	JSONData, err := json.MarshalIndent(userProile, "", "\t")

	CheckNilErr(err)

	jsonDataInStrForm := string(JSONData)

	fmt.Println(jsonDataInStrForm)

	file, _ := os.Create("./userProfile.json")

	io.WriteString(file, jsonDataInStrForm)

	defer file.Close()

	DecodeJSON()

}

func DecodeJSON() {
	fmt.Println("Decoding JSON data!")
	// ReadFile()

	var userData Profile

	file, err := ioutil.ReadFile("./userProfile.json")
	CheckNilErr(err)

	jsonDataFromFile := []byte(file)

	// isJSONValid := json.Valid(jsonDataFromFile)

	json.Unmarshal(jsonDataFromFile, &userData)

	fmt.Printf("%#v\n", userData)

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromFile, &myOnlineData)

	for _, val := range myOnlineData {
		fmt.Printf("%#v\n", val)
	}

	/*
		if isJSONValid {
			json.Unmarshal(isJSONValid, &userData)
			fmt.Printf("%#v\n", userData)

		} else {
			fmt.Println("JSON is not Valid.")
		}*/

}

type Profile struct {
	Name     string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"emailaddress"`
	Age      int    `json:"age"`
}

func structIterator([]reflect.StructField) {
	fields := reflect.VisibleFields(reflect.TypeOf(struct{ Profile }{}))
	for _, field := range fields {
		fmt.Printf("Key: %s\tType: %s\n", field.Name, field.Type)
	}
	// return fields
}

func ReadFile() {
	data, err := ioutil.ReadFile("./userProfile.json")
	CheckNilErr(err)
	// return data
	fmt.Printf("Type of data is %T\n", data)

}

func onlyInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input
}

func userInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return stoi(input)
}

func stoi(stringText string) int {
	stoi, err := strconv.Atoi(strings.TrimSpace(stringText))
	CheckNilErr(err)
	return stoi
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
