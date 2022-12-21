package main

import "fmt"

func main() {

	// var (
	// 	// number int
	// 	query int
	// )

	// fmt.Printf("Enter \n 1- Addition\n 2- Subtraction: ")
	// fmt.Scanf("%d\n", &query)
	// fmt.Printf("Enter radius: ")
	// fmt.Scanf("%d\n", &number)
	// printResult(10, 20, getFunction(query))
	// fmt.Println(addition(10, 20))
	// fmt.Println(subtraction(20, 10))

	// result, result1 := debug(addition(10, 20), subtraction(20, 10))
	// fmt.Println(result, result1)

	fmt.Println("---------- Calling first class function ----------")
	firtClassFunc(addition)
	firtClassFunc(subtraction)
	firtClassFunc(addition, subtraction, multiplication)
	firtClassFunc(division)

}

// Variadic function - passing function as an input
func firtClassFunc(calculateFunc ...func(int, int) int) {

	for _, val := range calculateFunc {
		fmt.Printf("%v\n", val(5, 2))
	}
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

func multiplication(num1, num2 int) int {
	return num1 * num2
	// fmt.Println(20 - 10)
	// return result
}

func division(num1, num2 int) int {
	return num1 / num2
	// fmt.Println(num1 / num2)
	// return result
}

// Working
func printResult(num1, num2 int, calculateFunc func(n1, n2 int) int) {
	result := calculateFunc(num1, num2)
	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Thank you!")
}

// Working
// func getFunction(query int) func(n1, n2 int) int {

// 	query_to_func := map[int]func(n1, n2 int) int{
// 		1: addition,
// 		2: subtraction,
// 	}
// 	return query_to_func[query]
// }
