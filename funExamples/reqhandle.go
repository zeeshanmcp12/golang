package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const acUrl = "https://lco.dev"

func main() {
	fmt.Println("Web request handling in golang!")

	// fmt.Println(acUrl)

	response, err := http.Get(acUrl)

	CheckNilErr(err)

	// fmt.Println(response)

	data, err := ioutil.ReadAll(response.Body)
	CheckNilErr(err)

	// fmt.Printf("Type of data is %T", string(data))

	// fmt.Println(string(data))
	defer response.Body.Close()

	storeRes(string(data))

}

func storeRes(data string) {
	resBody := data

	file, err := os.Create("./site_response.html")
	CheckNilErr(err)

	io.WriteString(file, resBody)
	defer file.Close()
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
