package main

import "fmt"

// func modify(y int) int {
// 	y += 15
// 	return y
// }

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

// func main() {
// 	s := "hello"
// 	var ptr *string = &s
// 	fmt.Println(s)
// 	*ptr += strings.ToUpper(s)
// 	fmt.Println(s)
// }

// Passing by value/reference to a function Lab
// Ex 1
// func modify(numbers ...int) {
// 	for i := range numbers {
// 		numbers[i] -= 5
// 	}
// }
// func main() {
// 	arr := []int{10, 20, 30}
// 	fmt.Println(arr)
// 	modify(arr...)
// 	fmt.Println(arr)
// }

// ------------------------------------------------
// Ex 2

// func modify(numbers [3]int) {
// 	for i := range numbers {
// 		numbers[i] -= 5
// 	}
// }
// func main() {
// 	arr := [3]int{10, 20, 30}
// 	fmt.Println(arr)
// 	modify(arr)
// 	fmt.Println(arr)
// }

// --------------------------------------------------
// Ex 3

// func modify(numbers *[3]int) {
// 	for i := range numbers {
// 		numbers[i] -= 5
// 	}
// }
// func main() {
// 	arr := [3]int{10, 20, 30}
// 	fmt.Println(arr)
// 	modify(arr)
// 	fmt.Println(arr)
// }

// -------------------------------------------------

// Ex 4

func modify(s map[string]int) {
	s["A"] = 100
}
func main() {
	ascii_codes := map[string]int{}
	ascii_codes["A"] = 65
	fmt.Println(ascii_codes)
	modify(ascii_codes)
	fmt.Println(ascii_codes)
}
