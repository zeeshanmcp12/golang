package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const siteUrl = "http://localhost:8000/"

func main() {
	fmt.Println("Web request with Get method!")
	GetRequest(siteUrl)
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

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
