package main

import "fmt"

/*
// Inner - Outer block
func main() {
	city := "ISB"

	{
		country := "Pakistan"
		fmt.Println("Country from the inner block", country)
		fmt.Println("City from inner block", city)
	}
	fmt.Println("City from outer block", city)
	// fmt.Println("Country from inner block", country)
}*/

// Local vs Global Variable
var globalVar string = "Global Variable"

func main() {

	// Local variable
	var name string = "Zeeshan"
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Global scoped variable: %v", globalVar)

}
