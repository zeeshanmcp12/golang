package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Following packages and methods are used to work with web verbs (get/post) etc
// http
// 	http.Get
// 	http.Post
// 	http.PostForm
// url
// 	url.Values
// ioutil
// 	ioutil.ReadAll
// strings
// 	strings.Builder
// 	strings.NewReader

func main() {
	fmt.Println("Web verbs i.e. get, post etc")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()

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

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname", "Muhammad")
	data.Add("lastname", "Zeeshan")
	data.Add("blog", "https://acloudtechie.com")

	response, err := http.PostForm(myUrl, data)
	checkNilErr(err)

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	// fmt.Println(string(content))
	fmt.Printf("Type of content: %T", content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
