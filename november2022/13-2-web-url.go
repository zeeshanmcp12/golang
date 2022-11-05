package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://example.com/contact"

func main() {
	fmt.Println("Working with web and url")

	response, err := http.Get(url)

	checkNilErr(err)

	fmt.Println(response.StatusCode)
	fmt.Printf("Type of response:%T\n", response.StatusCode)

	if response.StatusCode == 200 {
		fmt.Println("Site is working!")

		databyte, err := ioutil.ReadAll(response.Body)

		checkNilErr(err)

		fmt.Printf("Type -> %T\n", databyte)

	} else {
		fmt.Println("Site is not working!")
	}

	defer response.Body.Close()

	// getBody(url)

}

func getBody(url string) {
	response, err := http.Get(url)
	checkNilErr(err)

	databyte, err := ioutil.ReadAll(response.Body)
	checkNilErr(err)
	defer response.Body.Close()

	fmt.Println(string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
