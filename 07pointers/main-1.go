package main

import "fmt"

//
// - There are two operators
//       - & (amperson)
//       - * (asterik)
//     - & (amperson) gives us the address (where the variable is store in memory location)
//     - * (asterik) gives us the value stored at an address when we have the address

func main() {
	fmt.Println("Working with pointers in golang!")

	num1 := 7
	fmt.Println(num1)
	fmt.Printf("Type of num1 %T\n", num1)
	fmt.Printf("Address of num1 %v\n", &num1)

	addressTonum1 := &num1
	fmt.Println(*addressTonum1) // it will give us the value that is stored in &num1
	fmt.Printf("Type of addressTonum1 %T\n", addressTonum1)
	fmt.Printf("Address of addressTonum1 %v\n", &addressTonum1)
	fmt.Println(*&addressTonum1)

}
