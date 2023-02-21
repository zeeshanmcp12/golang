package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// func main() {

// 	result := addNum(10, 20)
// 	fmt.Println(result)

// 	greeter("Zeeshan")

// 	arr := []int{10, 20, 30}

// 	fmt.Println("Without modification", arr)

// 	// addOnlyFive(arr...)
// 	// fmt.Println("With modification by adding 5", arr)
// 	subOnlyFive(arr...)
// 	fmt.Println("With modification by subtracting 5", arr)

// 	// var name string = "Zeeshan"
// 	// for i, val := range name {
// 	// 	fmt.Printf("%v%c\n", i+1, val)
// 	// }
// }

func main() {
	fmt.Println("Todo app logic!")
	fmt.Printf("Add task: ")

	tasks := []string{}

	for true {
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		ftxt := strings.TrimSpace(input)

		if ftxt != "done" {
			if ftxt != "" {
				tasks = append(tasks, ftxt)
				fmt.Printf("Type done or add task: ")
			} else {
				fmt.Printf("Try again, type done or add task: ")
			}
			continue
		} else if ftxt == "done" {
			fmt.Println("Thank you for adding tasks: ")
			for i, val := range tasks {
				fmt.Printf("%v -> %v\n", i, val)
			}
		} else {
			fmt.Println("Invalid input, try again!")
		}
		break
	}
}

func addNum(num1, num2 int) int {
	sum := num1 + num2
	return sum
}

func greeter(text string) {
	fmt.Println("Hello", text)
}

func addOnlyFive(numbers ...int) {
	for i := range numbers {
		numbers[i] += 5
	}
}

func subOnlyFive(numbers ...int) {
	for i := range numbers {
		numbers[i] -= 5
	}
}
