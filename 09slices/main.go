package main

import "fmt"

func main() {
	fmt.Println("Learning about Slices in golang!")

	var sliceList = []string{"Apple", "Mango", "Banana"}
	fmt.Printf("Type of sliceList var is %T\n", sliceList)
	fmt.Printf("--------------------------------\n")
	fmt.Printf("sliceList before adding values\n")
	fmt.Println(sliceList)
	fmt.Printf("--------------------------------\n")
	fmt.Printf("sliceList after adding values\n")
	sliceList = append(sliceList, "Peach", "Orange")
	fmt.Println(sliceList)
	fmt.Printf("--------------------------------\n\n")
	fmt.Println("-----Slicing in Slice in golang------")
	sliceList = append(sliceList[1:])
	fmt.Println(sliceList)
	sliceList = append(sliceList[1:3])
	fmt.Println(sliceList)
	sliceList = append(sliceList[1:4])
	fmt.Println(sliceList)

}
