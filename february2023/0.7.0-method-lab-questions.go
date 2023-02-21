package main

import "fmt"

type Rectangle struct {
	length  int
	breadth int
}

func (r Rectangle) area() int {
	return r.length * r.breadth
}

func (r *Rectangle) incLength(n int) {
	for i := 0; i < n; i++ {
		// r.length += i
		r.length = r.length + i
		// Explanation below when n = 7
		/*
			0 -> 5
			1 -> 6
			2 -> 8
			3 -> 11
			4 -> 15
			5 -> 20
			6 -> 26
			7 -> this time loop will end because 7 is not less than 7 but equal
		*/
		// Explanation below when n = 6
		/*
			for i := 0; i < 6; i++ { r.length (5) = 5 + 0 }
			for i := 1; i < 6; i++ { r.length (6) = 5 + 1 }
			for i := 2; i < 6; i++ { r.length (8) = 6 + 2 }
			for i := 3; i < 6; i++ { r.length (11) = 8 + 3 }
			for i := 4; i < 6; i++ { r.length (15) = 11 + 4 }
			for i := 5; i < 6; i++ { r.length (20) = 15 + 5 }
			for i := 6; i < 6; i++ { Loop end because 6 is not less than 6 but equal }
		*/
		// fmt.Println(r.length)
	}
}

// func main() {
// 	r := Rectangle{breadth: 10, length: 5}
// 	fmt.Println(r.area())
// 	fmt.Printf("%+v\n", r)
// 	r.incLength(6)
// 	fmt.Println(r.area())
// 	fmt.Printf("%+v", r)
// }

// Next Question
// type Employee struct {
// 	eid int
// 	id  int
// }

// func main() {
// 	employees := make([]Employee, 5)
// 	for i := range employees {
// 		employees[i] = Employee{i, i + 10}
// 		fmt.Println(employees[i])
// 	}
// }

/* Output
{0 10}
{1 11}
{2 12}
{3 13}
{4 14}
*/

// Next Question
type Employee struct {
	eid int
	id  int
}

func (e Employee) get_id() int {
	return e.eid + 10
}

func main() {
	employees := make([]Employee, 5)
	for i := range employees {
		employees[i] = Employee{eid: i}
		// fmt.Printf("%+v\n", employees)
		employees[i].id = employees[i].get_id()
		// fmt.Printf("%+v\n", employees[i])
		fmt.Printf("%+v\n", employees[i])

	}
}
