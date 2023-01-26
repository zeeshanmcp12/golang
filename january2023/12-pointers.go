package main

import "fmt"

func modifyNumInMain(num int) {
	num += 20
}

func main() {
	fmt.Println("Practicing pointers in golang!")

	fmt.Println("Passing by value to function")
	var numInMain int = 10
	fmt.Println("Value of numInMain:", numInMain)
	modifyNumInMain(numInMain)
	fmt.Println("Value of numInMain after modification:", numInMain)

	fmt.Println("-------------------------")
	fmt.Println("Passing by reference to function")

}
