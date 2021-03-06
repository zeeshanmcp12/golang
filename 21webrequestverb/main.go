package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Making Get/Post/PostForm request to server!")
	// GetRequest()
	// PostRequest()
	PostFormRequest()
}

func GetRequest() {

	const localhost = "http://localhost:8000/get"

	response, err := http.Get(localhost)

	CheckNilErr(err)

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength) // In some cases, content length is important

	// One way of doing get request and print it
	content, err := ioutil.ReadAll(response.Body)
	CheckNilErr(err)
	fmt.Println(string(content))

	// Another way of doing get request by using Strings() package that has Builder functionality.
	// At this stage, I dont follow the reason to use this way but this is how it is.
	/*
		var responseString strings.Builder
		content, err := ioutil.ReadAll(response.Body)
		CheckNilErr(err)
		byteCount, _ := responseString.Write(content)
		fmt.Println(byteCount)
		fmt.Println(responseString.String())*/

	defer response.Body.Close()

}

func PostRequest() {
	const myUrl = "http://localhost:8000/post"
	// fake json payload
	requestBody := strings.NewReader(`
		{
			"name":"Zeeshan",
			"age":32,
			"website":"acloudtechie.com"
		}
	`)
	response, err := http.Post(myUrl, "application/json", requestBody)
	CheckNilErr(err)

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	CheckNilErr(err)
	fmt.Println(string(content))

}

func PostFormRequest() {
	myUrl := "http://localhost:8000/postform"

	data := url.Values{}

	data.Add("Name", "Zeeshan")
	data.Add("Age", "32")
	data.Add("Website", "https://acloudtechie.com")

	response, err := http.PostForm(myUrl, data)
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
