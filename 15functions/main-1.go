package main

import "fmt"

func main() {
	fmt.Println("Using functions in golang!")
	greeter()

	result, resMsg := addTwoNum(5, 9)
	fmt.Println("Combined two vars:", result, resMsg)

	proRes := variadicFunc(3, 5, 6, 9)
	fmt.Println("Pro result: ", proRes)

	table(2)

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

func table(numOne int) {
	for i := 0 + 1; i < 10+1; i++ {
		fmt.Printf("%d x %d = %d\n", numOne, i, numOne*i)

	}
}
