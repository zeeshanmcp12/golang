package main

import "fmt"

type Circle struct {
	area   float64
	radius float64
}

func main() {
	fmt.Println("Method in golang!")
	circle := Circle{radius: 5}
	fmt.Printf("%+v\n", circle)
	circle.areaOfCircle()
	fmt.Printf("%+v\n", circle)

	// fmt.Println(circle)
}

func (c *Circle) areaOfCircle() {
	PI := 3.14
	(*c).area = PI * c.radius * c.radius
}

// Value will not be modified because of not using pointer to an instance of Circle
// func (c Circle) areaOfCircle() {
// 	PI := 3.14
// 	c.area = PI * c.radius * c.radius
// }
