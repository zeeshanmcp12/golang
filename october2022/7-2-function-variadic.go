package main

import "fmt"

func main() {
	fmt.Println("Writing and calling variadic function in golang!")

	fmt.Println("Sum of two numbers:", variadicFunc(1, 2))
	fmt.Println("Sum of three numbers:", variadicFunc(1, 2, 3))
	fmt.Println("Sum of four numbers:", variadicFunc(1, 2, 3, 4))

	nums := []int{1, 2, 3, 4, 5, 6}
	addNumFromSlice := variadicFunc(nums...)
	fmt.Println("---------- Adding numbers from slice by using variadic function ----------")
	fmt.Println(addNumFromSlice)
}

func variadicFunc(numbers ...int) int {
	total := 0

	for _, val := range numbers {
		total += val
		// total = total + val
	}
	// return numbers[]
	return total

}
