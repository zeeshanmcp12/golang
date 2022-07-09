package main

import "fmt"

func main() {
	fmt.Println("Array in golang!")

	var list [4]string
	list[0] = "item1"
	list[1] = "item2"
	list[2] = "item3"

	fmt.Println("Elements in an array: ", len(list))
	fmt.Println("Total items in the list: ", list)

	var anotherList = [4]string{"item4", "item5", "item6", "item7"}
	fmt.Println(anotherList)

}
