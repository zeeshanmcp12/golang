package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://acloudtechie.com/blog?category=azure&name=test"

func main() {
	fmt.Println("Handling url")

	result, _ := url.Parse(myUrl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

}
