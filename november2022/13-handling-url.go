package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://acloudtechie.com/blog?category=azure&name=hello"

func main() {
	fmt.Println("Handling url in golang!")

	result, _ := url.Parse(myUrl)

	fmt.Println("URL Scheme i.e. http/https: ", result.Scheme)
	fmt.Println("Website host: ", result.Host)
	fmt.Println("API:", result.Path)
	fmt.Println("Query parameters: ", result.RawQuery)

	fmt.Printf("Type of RawQuery is: %T", result.RawQuery)

}
