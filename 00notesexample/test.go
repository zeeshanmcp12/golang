package main

import (
	"fmt"
	"math"
)

var (
	weight, height, bmi float64
)

func main() {
	fmt.Println("Enter your weight in kilograms : ")
	fmt.Scanf("%f", &weight)

	fmt.Println("Enter your height in meters : ")
	fmt.Scanf("%f", &height)

	fmt.Println("You have entered weight of : ", weight)
	fmt.Println("You have entered height of : ", height)

	bmi = weight / math.Pow(height, 2)

	fmt.Println("Your BMI is : ", bmi)
	fmt.Print("Your risk category is : ")

	if bmi < float64(18.5) {
		fmt.Println("Underweight")
	} else if bmi < 25 {
		fmt.Println("Normal weight")
	} else if bmi < 30 {
		fmt.Println("Overweight")
	} else {
		fmt.Println("Obese")
	}

	// calculate normal weight based on height and bmi = 25
	normalWeight := 25 * math.Pow(height, 2)
	delta := weight - normalWeight

	fmt.Printf("The normal weight for your height is : %0.2v kilograms.\n", normalWeight)

	if (delta > 0) && (bmi > 30) {
		fmt.Printf("You need to reduce %0.2v kilograms.\n", math.Abs(delta))
	}

	if (delta < 0) && (bmi < float64(18.5)) {
		fmt.Printf("You need to increase %0.2v kilograms.\n", math.Abs(delta))
	}

}
