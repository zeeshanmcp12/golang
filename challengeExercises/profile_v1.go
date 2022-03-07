package main

import (
	"bufio"
	"fmt"
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
	// onlyInput()
	password := strings.TrimSpace(onlyInput())

	fmt.Printf("Enter Email Address: ")
	// onlyInput()
	email := strings.TrimSpace(onlyInput())

	fmt.Printf("Enter your Age: ")
	agestoi := userInput()

	userProile := Profile{username, password, email, agestoi}

	value := reflect.ValueOf(userProile)

	typeOfStruct := value.Type()

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfStruct.Field(i).Name, value.Field(i).Interface())

	}

}

type Profile struct {
	Name     string `json:"username"`
	Password string `json:"-"`
	Email    string `json:"emailaddress"`
	Age      int    `json:"age"`
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
