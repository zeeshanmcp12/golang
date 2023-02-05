package main

import "fmt"

type Student struct {
	Name   string
	rollNo int
	marks  []int
}

func main() {

	// var studentData Student
	// fmt.Printf("%+v\n", studentData)
	// fmt.Printf("Address of studentData %v\n", &studentData)

	// newStudent := &studentData
	// fmt.Printf("Value of newStudent %v, and value at the address of newStudent %v", &newStudent, *newStudent)

	result := addNum(10, 20)
	fmt.Println(result)

	greeter("Zeeshan")

	arr := []int{10, 20, 30}

	fmt.Println("Without modification", arr)

	// addOnlyFive(arr...)
	// fmt.Println("With modification by adding 5", arr)
	subOnlyFive(arr...)
	fmt.Println("With modification by subtracting 5", arr)

	// var name string = "Zeeshan"
	// for i, val := range name {
	// 	fmt.Printf("%v%c\n", i+1, val)
	// }

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
