package main

import "fmt"

const (
	a = iota + 1
	b
	c
	d = iota + 2
	e
	f
	g = iota + 1
	h
	i
)

func main() {
	fmt.Println("iota in golang")

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", b)
	// fmt.Printf("%T\n", c)

}
