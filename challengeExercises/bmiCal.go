package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var weight, height, bmi float64

func main() {
	fmt.Println("BMI Calculator!")
	fmt.Printf("Enter your Weight: ")
	// weight := userInputToInt()
	// fmt.Scanf("%f", &weight)
	reader := bufio.NewReader(os.Stdin)
	weight, _ := reader.ReadString('\n')
	convWeight, err := strconv.ParseFloat(strings.TrimSpace(weight), 64)
	CheckNilErr(err)

	fmt.Printf("Enter your height: ")
	// height := userInputToInt()
	// reader := bufio.NewReader(os.Stdin)
	height, _ := reader.ReadString('\n')
	convHeight, err := strconv.ParseFloat(strings.TrimSpace(height), 64)
	CheckNilErr(err)
	// fmt.Scanf("%f", &height)

	// result := calculateBMI(float64(height), weight)
	bmiResult := convWeight / convHeight * 2

	fmt.Printf("Your BMI: %v", bmiResult)

}

// func calculateBMI(height int, weight float64) int {
// 	return height / float64(weight*2)

// }

func userInputToInt() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strToInt(input)
}

func strToInt(strText string) int {
	input, err := strconv.Atoi(strings.TrimSpace(strText))
	CheckNilErr(err)
	return input
}

func onlyString() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
