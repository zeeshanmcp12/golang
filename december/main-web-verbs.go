package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web verbs i.e. get, post etc")
	// PerformGetRequest()
	PerformPostRequest()

}

func PerformGetRequest() {
	const localhostUrl = "http://localhost:8000"

	response, err := http.Get(localhostUrl)
	checkNilErr(err)

	defer response.Body.Close()

	fmt.Println("Response status: ", response.Status)

	resBody, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(resBody))

	var responseBuilder strings.Builder
	stringBody, err := responseBuilder.Write(resBody)
	fmt.Println("Content length: ", stringBody)
	fmt.Println(responseBuilder.String())

}

func PerformPostRequest() {
	const myUrl = "http://localhost:8000/post"

	// fake json payload
	jsonData := strings.NewReader(`
		{
			"author":"zeeshan",
			"blog":"http://acloudtechie.com"
		}
	`)
	response, err := http.Post(myUrl, "application/json", jsonData)
	checkNilErr(err)
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
