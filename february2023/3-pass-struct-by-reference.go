package main

import "fmt"

func main() {
	fmt.Println("Passing struct by reference")

	circleProperties := Circle{radius: 10, area: 0}
	fmt.Printf("Area of a Circle -> %+v\n", circleProperties)
	calcAreaOfCircle(circleProperties)
	fmt.Printf("Area of a Circle calculating in function -> %+v\n", circleProperties)
	calcAreaByPassingRef(&circleProperties)
	fmt.Printf("Address of circleProperties %v\n", &circleProperties)
	fmt.Printf("Area of a Circle calculating in function by passing reference of a struct -> %+v\n", circleProperties)

}

type Circle struct {
	radius float64
	area   float64
}

func calcAreaOfCircle(c Circle) {
	var PI float64 = 3.14
	area := (PI * c.radius * c.radius)
	c.area = area
}

func calcAreaByPassingRef(c *Circle) {
	var PI float64 = 3.14
	area := (PI * c.radius * c.radius)
	(*c).area = area
}
