package main

import "fmt"

func main() {
	// debug()
	var (
		// number int
		query int
	)

	fmt.Printf("Enter \n 1- Addition\n 2- Subtraction: ")
	fmt.Scanf("%d\n", &query)
	// fmt.Printf("Enter radius: ")
	// fmt.Scanf("%d\n", &number)
	printResult(10, 20, getFunction(query))
	// addition()
	// subtraction()

}

func debug() {
	fmt.Println(20 + 10)
}

func printResult(num1, num2 int, calculateFunc func(n1, n2 int) int) {
	result := calculateFunc(num1, num2)
	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Thank you!")
}

func addition(num1, num2 int) int {
	return num1 + num2
	// fmt.Println(10 + 20)
	// return result
}

func subtraction(num1, num2 int) int {
	return num1 - num2
	// fmt.Println(20 - 10)
	// return result
}

func getFunction(query int) func(n1, n2 int) int {

	query_to_func := map[int]func(n1, n2 int) int{
		1: addition,
		2: subtraction,
	}
	return query_to_func[query]

}
