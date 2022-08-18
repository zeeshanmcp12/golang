package main

import "fmt"

func main() {
	fmt.Println("Using functions in golang!")
	greeter()

	result, resMsg := addTwoNum(5, 9)
	fmt.Println("Combined two vars:", result, resMsg)

	proRes := variadicFunc(3, 5, 6, 9)
	fmt.Println("Pro result: ", proRes)

}

func greeter() {
	fmt.Println("Hi, from golang!")
}

func addTwoNum(numOne, numTwo int) (string, int) {
	return "Sum: ", numOne + numTwo
}

func variadicFunc(values ...int) int {

	total := 0

	for _, val := range values {
		total += val
	}
	return total

}
