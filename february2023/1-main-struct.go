package main

import "fmt"

type Student struct {
	Name   string
	rollNo int
	marks  []int
}

func main() {

	var studentData Student
	fmt.Printf("%+v\n", studentData)

	studentData.Name = "Zeeshan"
	studentData.rollNo = 1001
	studentData.marks = []int{45, 50, 48}
	fmt.Printf("Address of studentData %v\n", &studentData)

	newStudent := &studentData
	fmt.Printf("Value at the address of newStudent %v\n", *newStudent)

	*newStudent = Student{"Abdullah", 1002, []int{50, 50, 48}}

	fmt.Printf("New Student added: %+v", studentData)

}
