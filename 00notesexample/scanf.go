package main

import (
	"fmt"
)

func main() {
	fmt.Println("Scanf function in golang")

	var weight float64
	var height float64

	fmt.Print("Enter your weight and height: ")
	fmt.Scanf("%f %f", &weight, &height)

	// fmt.Print("Enter your height: ")
	// fmt.Scanf("%.2f", &height)

	bmi := weight / height * 2

	fmt.Printf("Your Weight: %.1f\n", weight)
	fmt.Printf("Your Weight: %.1f\n", height)
	fmt.Printf("Your BMI: %.2f", bmi)

}
