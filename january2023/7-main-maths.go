package main

import "fmt"

func main() {
	fmt.Println("Print Table in golang!")

	var num int

	fmt.Printf("Enter any number: ")
	fmt.Scanf("%v", &num)

	for i := 1; i < 10+1; i++ {

		fmt.Printf("%v x %v = %v\n", num, i, num*i)

	}
}
