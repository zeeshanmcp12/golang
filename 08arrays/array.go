package main

import "fmt"

func main() {
	fmt.Println("Array in golang!")

	var list [4]string
	list[0] = "Apple"
	list[1] = "Potatoe"
	list[2] = "Orange"
	list[3] = "Banana"

	// fmt.Println("Elements in an array: ", len(list))
	// fmt.Println("Total items in the list: ", list)

	// var anotherList = [4]string{"item4", "item5", "item6", "item7"}
	// fmt.Println(anotherList)

	for i := 0; i < len(list); i++ {
		fmt.Printf("Item # %v is %v\n", i, list[i])

	}

}
