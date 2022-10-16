package main

import "fmt"

func main() {
	fmt.Println("Array in golang!")

	var newIteminArr string

	arr := [4]string{"Zeeshan", "32", "abc@go.dev"}

	// fmt.Println(len(arr))

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Index: %v, Value: %v\n", i, arr[i])

		if arr[i] == "" {
			fmt.Print("One item is empty, please add new: ")
			fmt.Scanf("%v", &newIteminArr)
			arr[i] = newIteminArr
			fmt.Printf("New Item added dynamically: %v\n", arr[i])
		}
	}

	fmt.Println("Updated list: ")

	for key, val := range arr {
		fmt.Printf("Index: %v, Value: %v\n", key+1, val)
	}
}
