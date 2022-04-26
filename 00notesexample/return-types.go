package main

import "fmt"

func main() {
	fmt.Println("Return types, Multiple, Named and Variadic in golang!")

	fmt.Println("Sum of 10 and 20 is", sumNumbers(10, 20))
	table(0, 1, 2, 3, 4, 5, 6) // It works without any "index out of range" error
	// table(1, 2, 3, 4, 5, 6) // It thwos "index out of range" error

}

func table(numbers ...int) {

	for _, val := range numbers {
		for i := 1; i < 10+1; i++ {
			fmt.Printf("%d x %d = %d\t", numbers[val], i, numbers[val]*i)
		}
	}
}

func sumNumbers(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}
