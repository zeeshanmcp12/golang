package main

import "fmt"

func main() {
	fmt.Println("Writing and calling variadic function in golang!")

	// fmt.Println(variadicFunc(2, 3, 4))
	variadicFunc(2, 3, 4)

	// nums := []int{1, 2, 3, 4, 5, 6}
	// addNumFromSlice := variadicFunc2(nums...)
	// fmt.Println("---------- Adding numbers from slice by using variadic function ----------")
	// fmt.Println(addNumFromSlice)
}

func variadicFunc(numbers ...int) {

	for i := 1; i < 10+1; i++ {
		fmt.Printf("%v x %v = %v\n", numbers, i, i*2)
	}
	// return i
}

func variadicFunc2(numbers ...int) int {
	total := 1

	for _, val := range numbers {
		total *= val
		// total = total + val
	}
	// return numbers[]
	return total

}
