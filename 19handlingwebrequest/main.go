package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://dev-id.matas.dk"

func main() {
	fmt.Println("Handling web request in golang!")

	response, err := http.Get(url)
	CheckNilErr(err)

	fmt.Printf("Type of response is %T\n", response)
	// Type of response is *http.Response which is a pointer
	// fmt.Println(response.Status)
	// fmt.Println(response.Request)

	// Close connection because it's caller's responsibility
	defer response.Body.Close()

	// Read the response
	// In this case, we only need to read the Body and not header etc
	data, err := ioutil.ReadAll(response.Body)
	CheckNilErr(err)

	// Convert the response into string
	fmt.Print(string(data))

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}

}
