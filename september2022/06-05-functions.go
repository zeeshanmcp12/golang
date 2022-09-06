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

	var (
		radius float64
		query  int
	)

	fmt.Printf("Enter radius of a circle: ")
	fmt.Scanf("%f\n", &radius)
	fmt.Println()
	fmt.Print("Select any number to calculate:\n1: Area\n2: Perimeter\n3: Diameter: ")
	fmt.Scanf("%v\n", &query)

	printResult(radius, getFunction(query))
}

func printResult(radius float64, calcFunction func(radius float64) float64) {
	result := calcFunction(radius)
	fmt.Println("Result: ", result)
	fmt.Println("Thank you!")

}

func getFunction(selection int) func(radius float64) float64 {
	selection_to_func := map[int]func(radius float64) float64{
		1: calcArea,
		2: calcPerimeter,
		3: calcDiameter,
	}
	return selection_to_func[selection]
}

func calcArea(radius float64) float64 {
	return 3.14 * radius * radius
}

func calcDiameter(radius float64) float64 {
	return 2 * 3.14 * radius
}

func calcPerimeter(radius float64) float64 {
	return 2 * radius
}
