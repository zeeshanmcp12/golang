package main

import (
	"fmt"
	"net/url"
)

// Handling (or working with) url
// url.Parse
// url -> is a package
// Parse is a method in url package

const myUrl = "https://acloudtechie.com/blog?category=azure&greeting=hello"

func main() {
	fmt.Println("Handling url in golang!")

	result, _ := url.Parse(myUrl)

	// fmt.Println("URL Scheme i.e. http/https: ", result.Scheme)
	// fmt.Println("Website host: ", result.Host)
	// fmt.Println("API:", result.Path)
	fmt.Println("Query parameters: ", result.RawQuery)

	fmt.Printf("Type of RawQuery is: %T\n", result.RawQuery)
	// at this stage we can only have query parameters in single line and we cannot manipulate as per our needs. i.e.
	// Query parameters:  category=azure&name=hello
	// We want to extract any of the query parameter and use accordingly.
	// Query() function can help us to achieve the scenario

	// Parsing query parameters
	qparams := result.Query()
	fmt.Printf("Type of qparams is %T\n", qparams)

	fmt.Println("Query parameter is:", qparams["category"])

	for j, val := range qparams {
		fmt.Printf("key -> %v, value -> %v\n", j, val)
	}

	// Constructing url

	newUrl := &url.URL{
		Scheme:   "https",
		Host:     "acloudtechie.com",
		Path:     "blog",
		RawQuery: "category=azure",
	}

	fmt.Println("Url: ", newUrl.String())

}
