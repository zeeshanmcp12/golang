package main

import "fmt"

func main() {
	fmt.Println("Scanf function in golang")

	var weight float64
	var height float64

	fmt.Print("Enter your weight: ")
	fmt.Scanf("%f\n", &weight)

	fmt.Print("Enter your height: ")
	fmt.Scanf("%f\n", &height)

	bmi := weight / height * 2

	fmt.Printf("Your Weight: %.1f\n", weight)
	fmt.Printf("Your Weight: %.1f\n", height)
	fmt.Printf("Your BMI: %.2f", bmi)

}

/*
func main() {
	fmt.Println("scanf in golang!")
	var weight float64
	var height float64

	fmt.Printf("Enter weight: ")
	count, err := fmt.Scanf("%f\n", &weight)
	fmt.Printf("Enter height: ")
	count1, err := fmt.Scanf("%f\n", &height)
	// count, err := fmt.Scanf("%f\n%f", &weight, &height)

	fmt.Println("Count", count)
	fmt.Println("Count", count1)
	fmt.Println("Error", err)
	fmt.Println("Weight", weight)
	fmt.Println("Height", height)
}*/

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
