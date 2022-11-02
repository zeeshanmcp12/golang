package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://example.com/"

func main() {
	fmt.Println("Handling web requests in golang!")

	getStatus(url)

	fmt.Printf("Type of getStatus function is %T", getStatus)

	// getBody(url)
}

func getStatus(url string) {
	responseStatus, err := http.Get(url)
	checkNilErr(err)

	fmt.Printf("Type of responseStatus is %T\n", responseStatus)
	defer responseStatus.Body.Close()

	status := responseStatus.StatusCode
	fmt.Printf("Type of status: %T\n", status)
	fmt.Println("Status code:", responseStatus.StatusCode)

}

func getBody(url string) {
	response, err := http.Get(url)
	checkNilErr(err)

	dataByte, err := ioutil.ReadAll(response.Body)
	content := string(dataByte)

	defer response.Body.Close()

	fmt.Printf("Type of response %T\n", response)
	fmt.Printf("Type of content %T\n", content)

	fmt.Println("Here is the body: ")
	fmt.Println(content)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
