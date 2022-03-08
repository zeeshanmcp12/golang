package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
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
