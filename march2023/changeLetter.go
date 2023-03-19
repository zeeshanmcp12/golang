package main

import "fmt"

func main() {
	fmt.Println("Replace Letter")
	a := "a"
	// b := "b"

	letters := []byte(a)

	fmt.Printf("%s -> %d -> %b -> %#U", letters, letters, letters, letters)

	// for i := 0; i < len(letters); i++ {
	// 	fmt.Printf("%s -> %d -> %b -> %#U")

	// }
}
