package main

import "fmt"

func main() {
	// debug()
	var (
		number int
		query  int
	)

	fmt.Printf("Enter \n 1- Area\n 2- Perimeter\n 3- Diameter: ")
	fmt.Scanf("%d\n", &query)
	fmt.Printf("Enter radius: ")
	fmt.Scanf("%d\n", &number)
	printResult(number, getFunction(query))
	addition()
	subtraction()

}

func debug() {
	fmt.Println(20 + 10)
}

func printResult(number int, calculateFunc func(n int) int) {
	result := calculateFunc(number)
	fmt.Printf("Result: %v\n", result)
	fmt.Printf("Thank you!")
}

func addition() {
	fmt.Println(10 + 20)
	// return result
}

func subtraction() {
	fmt.Println(20 - 10)
	// return result
}

func getFunction(query int) func() {

	query_to_func := map[int]func(){
		1: addition,
		2: subtraction,
	}
	return query_to_func[query]

}
