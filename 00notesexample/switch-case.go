package main

import "fmt"

func main() {
	fmt.Println("Switch Case in golang!")

	var number int = 1

	switch number {
	case 10:
		fmt.Println("It is 10")
	case 1:
		fmt.Println("I found 1")
		fallthrough
	case 2:
		fmt.Println("What to do next")
	default:
		fmt.Println("Sorry, 1 is lost!")
	}
}
