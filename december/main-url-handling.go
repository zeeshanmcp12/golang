package main

import (
	"fmt"
	"net/url"
)

// const myUrl = "https://example.com"
const myUrl = "https://acloudtechie.com/blog?category=azure&author=zeeshan"

func main() {
	fmt.Println("Handling url in golang!")

	parsedUrl, err := url.Parse(myUrl)
	checkNilErr(err)

	fmt.Println(parsedUrl.Scheme)
	fmt.Println(parsedUrl.Host)
	fmt.Println(parsedUrl.Path)
	fmt.Println(parsedUrl.RawQuery)

	result := parsedUrl.Query()
	fmt.Println(result)

	for i, val := range result {
		fmt.Printf("%v = %v\n", i, val)
	}

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
