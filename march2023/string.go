package main

import "fmt"

func main() {
	fmt.Println("Using `raw string literal`")

	multilineString := `"Multi line
	string"`
	fmt.Println(multilineString)
}
