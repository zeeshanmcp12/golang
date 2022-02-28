package main

import "fmt"

func main() {
	fmt.Println("defer in golang!")
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	deferFun()

}

// Hello, 0,1,2,3,4,

// Initially the defer will add values to stack
// Now the lines should be emptied up (this will )
func deferFun() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)

	}
}
