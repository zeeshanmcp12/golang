package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Handling Web request/response in golang")

	response, err := http.Get(url)
	CheckNilErr(err)

	fmt.Printf("Type of url is %T", response)

	databytes, err := ioutil.ReadAll(response.Body)
	CheckNilErr(err)
	fmt.Println(string(databytes))

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}

}
