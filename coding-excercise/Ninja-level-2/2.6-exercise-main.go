package main

import "fmt"

const (
	a = 2017 + iota
	b = 2017 + iota
	c = 2017 + iota
	d = 2017 + iota
	e
)

func main() {
	fmt.Println("Exercise 6 - Ninja level 2")

	fmt.Println(a, b, c, d)

	for i := 0; i < 10; i++ {
		fmt.Println(e + i)

	}
}
