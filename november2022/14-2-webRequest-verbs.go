package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const siteUrl = "http://localhost:8000"

func main() {
	fmt.Println("Working with web request verbs!")

	GetRequest()
	PostRequest()

}

func GetRequest() {
	const siteUrl = "http://localhost:8000"

	response, err := http.Get(siteUrl)
	CheckNilErr(err)

	requestBody, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(requestBody))

	defer response.Body.Close()
}

func PostRequest() {
	const siteUrl = "http://localhost:8000/post"

	jsonPayload := strings.NewReader(`
	{
		"name":"zeeshan"
	}
	`)

	response, err := http.Post(siteUrl, "application/json", jsonPayload)
	CheckNilErr(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("----- Data Processed -----")
	fmt.Println(string(content))

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
