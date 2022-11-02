package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Handling web requests
// - First we need to define url (may be constant in global scope)
// - Package -> http
// - Method -> Get, Post etc
// - We must need to close the response once we are finished
// - Close() -> response.Body.Close()
// - Read response whatever we received after requesting web (or website)
// - ioutil.ReadAll()

// const url = "https://acloudtechie.com"
const url = "https://google.com"

func main() {
	fmt.Println("Handling web request in golang!")
	// response, err := http.Get(url)

	// checkNilErr(err)

	// fmt.Printf("Type of response is %T\n", response)
	// defer response.Body.Close()

	// data, err := ioutil.ReadAll(response.Body)

	// content := string(data)

	// fmt.Println(content)

	// Check website status
	getStatus(url)

	// Print Body of web page
	getBody(url)

}

func getStatus(url string) {
	response, err := http.Get(url)
	checkNilErr(err)

	status := response.Status
	fmt.Printf("Type of status is %T\n", status)

	fmt.Println("Http status code is: ", status)
	defer response.Body.Close()
}

func getBody(url string) {

	response, err := http.Get(url)

	checkNilErr(err)

	fmt.Printf("Type of response is %T\n", response)
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)

	content := string(data)

	fmt.Println(content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
