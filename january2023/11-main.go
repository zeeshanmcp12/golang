package main

import (
	"fmt"
	"strings"
)

func modify(y int) int {
	y += 15
	return y
}

// func main() {
// 	y := 20
// 	modify(y)
// 	fmt.Println(y)
// }

// // Why 20

// func main() {
// 	y := [3]int{10, 20, 30}
// 	fmt.Printf("%v \n", y)
// 	(*&y)[0] = 100
// 	fmt.Printf("%v \n", y)
// }

// func main() {
// 	var y int
// 	var ptr *int = &y
// 	fmt.Println(y)
// 	fmt.Println(*ptr)
// }

func main() {
	s := "hello"
	var ptr *string = &s
	fmt.Println(s)
	*ptr += strings.ToUpper(s)
	fmt.Println(s)
}
