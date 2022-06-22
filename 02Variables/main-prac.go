package main

import "fmt"

const LoginName string = "Zeeshan"

func main() {
	fmt.Println("Variables")

	/*
		var username string = "gopher"
		fmt.Printf("Type of username is %T and value is %s", username, username)
		fmt.Println()

		var isLoggedHours bool = true
		fmt.Printf("Type of isLoggedHours var is %T and value is %t", isLoggedHours, isLoggedHours)
		fmt.Println()

		var amount int = 200
		fmt.Printf("Type of variable amount is %T and value is %d", amount, amount)
		fmt.Println()

		var floatingNum float32 = 3.14
		fmt.Printf("Type of variable floatingNum is %T and value is %.2f", floatingNum, floatingNum)
		fmt.Println()

		fmt.Printf("Type of LoginName is %T and value is %s\n", LoginName, LoginName)*/

	var weight float32
	var height float32

	fmt.Print("Enter your weight: ")
	fmt.Scanf("%f\n", &weight)

	fmt.Print("Enter your height: ")
	fmt.Scanf("%f\n", &height)

	var bmi float32 = weight / height * 2

	fmt.Println("Weight: ", weight)
	fmt.Println("Height: ", height)
	fmt.Printf("Your BMI is %.2f\n", bmi)

	if bmi > 18.5 && bmi < 25 {
		fmt.Println("You weight consider as Healthy!")

	} else if bmi > 25 {
		fmt.Println("You are overweight. Do workout everyday!")

	} else if bmi < 18.5 {
		fmt.Println("You are underweight. Gain some weight!")

	} else if bmi == 25 {
		fmt.Println("You are at borderline!")

	} else {
		fmt.Println("Invalid!")
	}

}
