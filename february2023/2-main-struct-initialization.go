package main

import "fmt"

type Student struct {
	name   string
	rollNo int
	marks  []int
	// grades map[string]int
}

func main() {
	fmt.Println("Struct Initialization in golang!")
	// First method
	var st Student
	fmt.Printf("Data in struct: %+v\n", st)

	// Second method is using 'new' keyword
	studentStruct := new(Student)
	fmt.Printf("Second method shows struct as pointer ->\nData in struct: %+v\n", studentStruct)

	struct3rdMethod := Student{
		name:   "Abdullah",
		rollNo: 1001,
		marks:  []int{35, 40, 50},
	}
	fmt.Printf("Data in struct using 3rd Method -> %+v\n", struct3rdMethod)

	// We can also omit the field name in following method:
	struct4rdMethod := Student{"Abdullah", 1002, []int{40, 45, 50}}
	fmt.Printf("Data in struct using 4th Method -> %+v\n", struct4rdMethod)

}
