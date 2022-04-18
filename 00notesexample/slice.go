package main

import "fmt"

func main() {
	fmt.Println("Slice in golang")

	var arr = [5]int{10, 20, 30, 40, 50}
	slice := arr[:3]
	fmt.Println("Before updating value")
	fmt.Println(arr)   // output: [10,20,30,40,50]
	fmt.Println(slice) // output: [10,20,30]

	slice[1] = 35
	fmt.Println("After updating value")
	fmt.Println(arr)   // output: [10,35,30,40,50]
	fmt.Println(slice) // output: [10,35,30]

	// copy elements from one slice to another
	destSlice := make([]int, 3)
	copiedElem := copy(destSlice, slice)

	fmt.Println("Elements copied from slice to destSlice: ", destSlice)
	fmt.Println("Number of elements copied: ", copiedElem)

}
