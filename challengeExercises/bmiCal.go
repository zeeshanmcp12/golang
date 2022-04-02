package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("BMI Calculator!")
	fmt.Printf("Enter your Weight: ")
	weight := userInputToInt()

	fmt.Printf("Enter your height: ")
	// height := userInputToInt()
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	height, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	CheckNilErr(err)

	// result := calculateBMI(float64(height), weight)
	result := weight / int(height) * 2

	fmt.Printf("Your BMI: %v", result)

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
