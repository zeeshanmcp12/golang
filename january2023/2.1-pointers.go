package main

import "fmt"

func main() {
	fmt.Println("Pointers in golang!")

	var num1, num2 = 43, 2701
	fmt.Printf("Value of num1 -> %v\nValue of num2 -> %v", num1, num2)
	fmt.Println()
	fmt.Printf("Address of num1 -> %v\nAddress of num2 -> %v", &num1, &num2)
	fmt.Println()
	p := &num1
	fmt.Printf("Value at the address of num1 -> %v\n", *p)
	fmt.Println("Let's change the value of num1 through p")
	*p = 50
	fmt.Printf("Now, the value at the address of num1 -> %v\n", num1)
	p = &num2
	fmt.Printf("Value at the address of num2 -> %v\n", *p)
	fmt.Println("Let's change the value of num2 through p")
	*p = *p / 73
	fmt.Printf("Now, the value at the address of num2 -> %v...\nBecause, 2701 / 73 = %v", num2, num2)

}
