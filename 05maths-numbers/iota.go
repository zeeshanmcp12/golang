package main

import "fmt"

const (
	a = iota + 1
	b
	c
)

const (
	x = 2 * iota // The value of iota starts at 0 so 2 * 0 = 0
	y            // iota increments the value to 1 to 2 * 1 = 2
	z            // iota increments the value to 1 to 2 * 2 = 4
)

func main() {
	fmt.Println("iota in golang")

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
