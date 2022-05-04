package main

import "fmt"

// Ask user to calculate the area of a circle
// Ask user to calculate the perimeter of a circle
// Ask user to calculate the diameter of a circle

func main() {
	var (
		selection int
		radius    float64
	)
	PI := 3.14
	fmt.Println("Calculate the properties of a Circle!")
	fmt.Printf("Enter the radius of a circle: ")
	fmt.Scanf("%f\n", &radius)

	fmt.Printf("Select: \n1: Area of a Circle \n2: Perimeter of a Circle \n3: Diameter of a Circle\n")
	fmt.Scanf("%d\n", &selection)

	if selection == 1 {
		area := PI * radius * radius
		fmt.Println("The Area of a circle is", area)

	} else if selection == 2 {
		perimeter := 2 * PI * radius
		fmt.Println("The Perimeter of a circle is", perimeter)

	} else if selection == 3 {
		diameter := 2 * radius
		fmt.Println("The Diameter of a circle is", diameter)

	} else {
		fmt.Println("Invalid Input!")
	}
}
