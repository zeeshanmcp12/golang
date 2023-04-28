package main

import "fmt"

func main() {
	fmt.Println("Dynamic or Static-Type")

	num := 5
	fmt.Println(num)
	fmt.Printf("%T\n", num)
	num = "five"
	fmt.Println(num)
	fmt.Printf("%T", num)
}
