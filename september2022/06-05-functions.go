package main

import "fmt"

// Formulas:
// Calculate the area
// 3.14 * radius * radius

// Calculate the diameter
// 2 * 3.14 * radius

// Calculate the perimeter
// 2 * radius

func main() {
	fmt.Println("Using High order function in golang!")
	fmt.Println("Area of the circle: ", calculateArea(9.1))
}

func calculateArea(radius float64) float64 {
	result := 3.14 * radius * radius
	return result
}
