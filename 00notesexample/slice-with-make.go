package main

import "fmt"

func main() {
	fmt.Println("Declaring slice with make function!")
	makeSlice := make([]int, 5, 8)

	fmt.Println(makeSlice)      // output [0 0 0 0 0] -> because we didn't initialize slice with values so the 0 is default value for nil
	fmt.Println(len(makeSlice)) // 5
	fmt.Println(cap(makeSlice)) // 8
}
