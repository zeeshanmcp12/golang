package main

import "fmt"

func main() {
	city := "ISB"

	{
		country := "Pakistan"
		fmt.Println("Country from the inner block", country)
		fmt.Println("City from inner block", city)
	}
	fmt.Println("City from outer block", city)
	// fmt.Println("Country from inner block", country)
}
