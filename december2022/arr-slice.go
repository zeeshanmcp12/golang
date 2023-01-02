package main

import "fmt"

func main() {
	fmt.Println("Slice revisit")

	num := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	// fmt.Println(num)

	// for i, val := range num {
	// 	fmt.Printf("%d'th place is %v\n", i, val)
	// }

	newSlice := num[2:5]
	fmt.Println("Before slicing: ", num)
	// fmt.Printf("%v", newSlice)
	for i := range newSlice {
		newSlice[i]++
	}
	fmt.Println("After slicing: ", num)
	// fmt.Printf("%v", newSlice)
}
