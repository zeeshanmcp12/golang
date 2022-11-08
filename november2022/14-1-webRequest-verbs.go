package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const siteUrl = "http://localhost:8000/post"

func main() {
	fmt.Println("Web request with Get method!")
	// GetRequest(siteUrl)
	PostRequest(siteUrl)
}

func GetRequest(siteUrl string) {

	response, err := http.Get(siteUrl)
	CheckNilErr(err)

	fmt.Println("Site status: ", response.Status)

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	// One way of wrapping up response body into string format
	// fmt.Println(string(content))

	// Another way of wrapping up response body into string format is using string builder.
	var responseBuilder strings.Builder
	responseBuilder.Write(content)
	fmt.Println(responseBuilder.String())

}

func PostRequest(url string) {

	// Create fake json payload
	request := strings.NewReader(`
	{
		"name":"zeeshan",
		"age":32,
		"email":"admin@acloudtechie.com"
	}
	`)

	// http -> package
	// Post -> method that accepts three arguments:
	// 			url, content format (for example "application/json"), request body or json payload.
	response, err := http.Post(url, "application/json", request)
	CheckNilErr(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
