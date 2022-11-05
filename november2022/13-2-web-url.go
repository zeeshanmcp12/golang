package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://example.com"

func main() {
	fmt.Println("Working with web and url")

	response, err := http.Get(url)

	checkNilErr(err)

	fmt.Println(response.StatusCode)
	fmt.Printf("Type of response:%T\n", response.StatusCode)

	file := "./13-2-webpage.txt"

	if response.StatusCode == 200 {
		fmt.Println("Site is working!")

		databyte, err := ioutil.ReadAll(response.Body)

		checkNilErr(err)

		// fmt.Printf("Type -> %T\n", string(databyte))
		// fmt.Println("Go find the content", string(databyte))

		content := string(databyte)

		filename, _ := os.Create(file)

		io.WriteString(filename, content)
		defer filename.Close()

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
