package main

import "fmt"

func main() {
	fmt.Println("Print Table")
	for d := 1; d < 10+1; d++ {
		fmt.Printf("2 x %v = %v\n", d, d*2)
	}
}
