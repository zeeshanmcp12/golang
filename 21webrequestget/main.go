package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Making Get request to server!")
	GetRequest()
}

func GetRequest() {

	const localhost = "http://localhost:8000/get"

	response, err := http.Get(localhost)

	CheckNilErr(err)

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)
	CheckNilErr(err)

	fmt.Println(string(content))

	defer response.Body.Close()

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
