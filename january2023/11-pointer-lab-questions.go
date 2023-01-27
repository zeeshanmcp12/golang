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
// Because above we are using `Pass by value to function` which does not change/modify the actual value because it's being copied to another location

// func main() {
// 	y := [3]int{10, 20, 30}
// 	fmt.Printf("%v \n", y)
// 	(*&y)[0] = 100
// 	fmt.Printf("%v \n", y)
// }
// In above and below example, it changes the value at 0th index to 100

// func main() {
// 	y := [3]int{10, 20, 30}
// 	fmt.Printf("%v \n", y)
// 	y[0] = 100
// 	fmt.Printf("%v \n", y)
// }

// func main() {
// 	var y int
// 	var ptr *int = &y
// 	fmt.Println(y)
// 	fmt.Println(*ptr)
// }

// Output
// 0
// 0
// Because y variable is declared with no value, and ptr contains the value at the address of y that's why it's output is 0

// func main() {
// 	s := "hello"
// 	var ptr *string = &s
// 	// var s_p string
// 	fmt.Println(s)
// 	fmt.Println(*ptr)
// 	*ptr += strings.ToUpper(s)
// 	fmt.Println(s)
// }

// Output
// hello
// helloHello
// hello is printed at line 53
// helloHELLO -> *ptr contains the value at the address of s which is hello (with small letters)
// HELLO -> strings.ToUpper(s) contains this capital HELLO.

// Passing by value/reference to a function Lab
// Ex 1
func modify(numbers ...int) {
	// for i := range numbers {
	// 	numbers[i] -= 5
	// }

	// Another way of writing loop
	for i := 0; i < len(numbers); i++ {
		// numbers[i] -= 5
		numbers[i] = numbers[i] - 5
	}
}
func main() {
	arr := []int{10, 20, 30}
	fmt.Println(arr)
	modify(arr...)
	fmt.Println(arr)
}

// Output
// [10 20 30]
// [5 15 25]
// Slice and maps are "Passed by reference to function" by default

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

// func modify(s map[string]int) {
// 	s["A"] = 100
// }
// func main() {
// 	ascii_codes := map[string]int{}
// 	ascii_codes["A"] = 65
// 	fmt.Println(ascii_codes)
// 	modify(ascii_codes)
// 	fmt.Println(ascii_codes)
// }
