package main

import "fmt"

func main() {
	fmt.Println("Capacity in slice/array")
	var arr = [10]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	mySlice := arr[1:8]
	fmt.Println(cap(arr))
	fmt.Println(cap(mySlice))
}
