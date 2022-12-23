package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Web verbs i.e. get, post etc")
	PerformGetRequest()

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

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
