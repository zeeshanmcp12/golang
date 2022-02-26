package main

import "fmt"

func main() {
	fmt.Println("Practice about functions in golang")

	// writing function inside function
	output := func() {
		fmt.Println("function call from inside of function")
	}
	output()

}
