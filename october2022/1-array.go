package main

import "fmt"

func main() {
	fmt.Println("Array in golang!")

	arr := [4]string{"Zeeshan", "32", "abc@go.dev"}

	// fmt.Println(len(arr))

	for i := 0; i < len(arr); i++ {
		fmt.Printf("Index: %v, Value: %v\n", i, arr[i])

		if arr[i] == "" {
			arr[i] = "Registered"
			fmt.Printf("New Item added dynamically: %v\n", arr[i])
		}
	}

	fmt.Println("Updated list: ")

	for key, val := range arr {
		fmt.Printf("Index: %v, Value: %v\n", key+1, val)
	}
}
