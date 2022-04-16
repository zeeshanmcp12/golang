package main

import "fmt"

func main() {
	fmt.Println("Slice in golang with example!")

	var arr []int = []int{1, 2, 4, 5, 67, 4, 7, 7}

	for index, element := range arr {
		fmt.Printf("%d: %d\n", index, element)
		// fmt.Println("Simple slice above")
		for j := index + 1; j < len(arr); j++ {
			element2 := arr[j]
			if element2 == element {
				// fmt.Println("Duplicate found below!")
				fmt.Println(element)
			}
		}
	}
}
