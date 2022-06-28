package main

import "fmt"

func main() {

	fmt.Println("Loops in golang!")
	var tableNum int

	fmt.Print("Enter any number: ")
	fmt.Scanf("%v\n", &tableNum)

	for i := 1; i < 10+1; i++ {
		fmt.Printf("%v x %v = %v\n", tableNum, i, i*tableNum)

	}
}
