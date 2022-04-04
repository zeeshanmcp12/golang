package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Find data type!")

	var number int
	var text string
	var isTrue bool
	var PI float64

	fmt.Printf("Enter any number: ")
	fmt.Scanf("%d\n", &number)

	fmt.Printf("Enter any text: ")
	fmt.Scanf("%s\n", &text)

	fmt.Printf("Is it true: ")
	fmt.Scanf("%t\n", &isTrue)

	fmt.Printf("Enter PI value: ")
	fmt.Scanf("%f\n", &PI)

	fmt.Printf("Type: %v and Value %v\n", reflect.TypeOf(number), number)
	fmt.Printf("Type: %v and Value %v\n", reflect.TypeOf(text), text)
	fmt.Printf("Type: %v and Value %v\n", reflect.TypeOf(isTrue), isTrue)
	fmt.Printf("Type: %v and Value %v\n", reflect.TypeOf(PI), PI)

}
