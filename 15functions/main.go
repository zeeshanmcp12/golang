package main

import "fmt"

func add(numOne, numTwo int) {
	fmt.Println(numOne + numTwo)

}

func main() {
	fmt.Println("functions in golang")
	greetings()
	add(5, 5)
}

func greetings() {
	fmt.Println("Printing from outside function")
}
