package main

import "fmt"

func main() {
	fmt.Println("Switch Case in golang!")

	var number int = 100

	switch number {
	case 100:
		fmt.Println("It is 100")
	case 90:
		fmt.Println("It is 90")
		// fallthrough
	case 2:
		fmt.Println("What to do next")
	default:
		fmt.Println("Sorry, 100 is lost!")
	}
}
