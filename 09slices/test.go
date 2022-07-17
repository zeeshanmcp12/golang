package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("testing scanf function")

	// var text string

	// fmt.Print("Enter any text: ")
	// // fmt.Scanf("%s\n", &text)

	// count, err := fmt.Scanf("%q\n", &text)

	// // if err != nil {
	// // 	panic(err)
	// // }

	// fmt.Println("Given input: ", text)
	// fmt.Println("Count: ", count)
	// fmt.Println("Error: ", err)

	// var items = []string{}
	// fmt.Printf("Type of items: %T\n", items)

	// items = append(items, "Apple", "Orange", "Peach", "Banana", "Mango", "Watermelon")

	// for i, item := range items {
	// 	fmt.Printf("%v item: %s\n", i, item)
	// }

	// var indexNum int = 4

	// items = append(items[1 : indexNum+1])
	// fmt.Println("item list updated after slice: ", items)

	alphabets := make([]string, 4)
	fmt.Printf("Type of alphabets var is %T", alphabets)

	alphabets = append(alphabets, "E", "C", "A", "B", "D")
	fmt.Println(alphabets)

	sort.Strings(alphabets)
	fmt.Println("Strings are sorted: ", alphabets)
	fmt.Println(sort.StringsAreSorted(alphabets))

}
