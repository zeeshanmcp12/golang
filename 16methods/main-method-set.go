package main

import "fmt"

type Student struct {
	name   string
	grades []int
}

func main() {
	fmt.Println("Method sets in golang!")
	student := Student{"Abdullah", []int{90, 80, 95}}
	// fmt.Printf("Student Data -> %+v\n", student)
	student.displayStudentName()
	// student.studentPercentage()

	fmt.Printf("Percentage -> %.2f%%", student.studentPercentage())
}

func (s Student) displayStudentName() {
	fmt.Println("Student Name -> ", s.name)
}

func (s Student) studentPercentage() float64 {
	sum := 0
	for _, val := range s.grades {
		sum += val
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}
