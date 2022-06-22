package main

import (
	"fmt"
)

/*
func main() {
	fmt.Println("Taking user input using Scanf function!")

	var weight float32
	var height float32

	fmt.Print("Enter your weight: ")
	fmt.Scanf("%f\n", &weight)

	fmt.Print("Enter your height: ")
	fmt.Scanf("%f\n", &height)

	fmt.Printf("Your weight is %.2f and height is %.2f", weight, height)
}*/

func main() {
	fmt.Println("Using Scanf function!")

	var weight float32
	var height float32

	fmt.Print("Enter your weight: ")
	count, err := fmt.Scanf("%f\n", &weight)

	fmt.Print("Enter your height: ")
	count1, err := fmt.Scanf("%f\n", &height)

	fmt.Println("Count: ", count)
	fmt.Println("Err: ", err)
	fmt.Println("Count1: ", count1)
	fmt.Println("Err: ", err)

	fmt.Println("Weight: ", weight)
	fmt.Println("Height: ", height)

	// fmt.Printf("Type of variables are %T", (reflect.TypeOf(weight)))

	if err != nil {
		panic(err)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
