package main

import (
	"fmt"
	"net/url"
)

const acloudtechie string = "https://acloudtechie.com:8443/blog?category=azure&tags=azuredevops"

func main() {
	fmt.Println("URL handling in golang!")

	fmt.Println("URL is", acloudtechie)
	// fmt.Printf("Type of url is %T\n", acloudtechie)

	// Parsing in golang
	// Parsing means we want to extract some information from the url (in context meaning)
	parsedUrl, err := url.Parse(acloudtechie)

	CheckNilErr(err)

	// fmt.Println("Protocol is", parsedUrl.Scheme)
	// fmt.Println("Host is", parsedUrl.Host)
	// fmt.Println("Port is", parsedUrl.Port()) // Port is a method when we use url.Parse
	// fmt.Println("Path is", parsedUrl.Path)
	fmt.Println("Query parameters are", parsedUrl.RawQuery)

	queryParams := parsedUrl.Query()
	// '.Query()' function holds / parse all the query parameters from url

	// fmt.Printf("Type of queryParams %T\n", queryParams)
	// fmt.Printf("Value of category is %v\n", queryParams["category"])
	// fmt.Printf("Value of tags is %v", queryParams["tags"])

	for i, val := range queryParams {
		fmt.Printf("Key is %v Value is %v\n", i, val)
		// i++
	}

	constructUrl := &url.URL{ // dont forget to pass the reference of url here and not the copy.
		Scheme: "https",
		Host:   "acloudtechie.com",
		Path:   "/azure",
	}

	homePage := constructUrl.String()
	fmt.Println(homePage)
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
