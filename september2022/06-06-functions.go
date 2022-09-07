package main

import "fmt"

func main() {
	fmt.Println("Using high order function in golang!")

	var (
		query  int
		radius float64
	)

	fmt.Print("Enter radius of a circle: ")
	fmt.Scanf("%v\n", &radius)

	fmt.Print("Select any number to calculate:\n1: Area\n2:Diameter\n3:Perimeter: ")
	fmt.Scanf("%v\n", &query)

	printResult(query, getFunction(radius))
}

func printResult(radius float64, calcFunction func(radius float64) float64) {

}

// Formulas:
// Calculate the area
// 3.14 * radius * radius

// Calculate the diameter
// 2 * 3.14 * radius

// Calculate the perimeter
// 2 * radius

func calcArea(radius float64) float64 {
	return 3.14 * radius * radius
}

func calcDiameter(radius float64) float64 {
	return 2 * 3.14 * radius
}

func calcPerimeter(radius float64) float64 {
	return 2 * radius
}
