package main

import "fmt"

func main() {
	fmt.Println("Revisiting some concepts")
	var number int

	fmt.Print("Enter any number: ")
	fmt.Scanf("%v", &number)

	// fmt.Println(number)

	for i := 1; i < 10+1; i++ {
		fmt.Printf("%v x %v = %v\n", number, i, number*i)

	}

}
