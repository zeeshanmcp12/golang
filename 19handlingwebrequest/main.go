package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://dev-id.matas.dk"

// This is about, How we can make a request to a URL

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

	// successReq := 200

	/*	if successReq == response.StatusCode {
			fmt.Printf("%v returning %v status ", url, response.Status)

		} else {
			fmt.Println("There is an error")
		}*/

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
