package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

	fmt.Println(string(content))

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
