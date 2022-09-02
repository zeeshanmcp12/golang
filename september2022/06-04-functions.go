package main

import (
	"fmt"
)

func main() {
	fmt.Println("Using high order function!")

	var radius float64
	var selection int

	// reader := bufio.NewReader(os.Stdin)
	// reader.ReadString('\n')

	fmt.Printf("Enter radius: ")
	fmt.Scanf("%f\n", &radius)

	fmt.Printf("Select option: \n1[Area] \n2[Diameter] \n3[Perimeter]: ")
	fmt.Scanf("%d\n", &selection)

	fmt.Println()

	fmt.Printf("Type is: %T, Value is: %v\n", radius, radius)
	fmt.Printf("Type is: %T, Value is: %v\n", selection, selection)
}

// Formulas:
// Calculate the area
// 3.14 * radius * radius

// Calculate the diameter
// 2 * 3.14 * radius

// Calculate the perimeter
// 2 * radius
