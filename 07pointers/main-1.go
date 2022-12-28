package main

import "fmt"

//
// - There are two operators
//       - & (amperson)
//       - * (asterik)
//     - & (amperson) gives us the address (where the variable is store in memory location)
//     - * (asterik) gives us the value stored at an address when we have the address

// func main() {
// 	fmt.Println("Working with pointers in golang!")

// 	num1 := 7
// 	fmt.Println(num1)
// 	fmt.Printf("Type of num1 %T\n", num1)
// 	fmt.Printf("Address of num1 %v\n", &num1)

// 	addressTonum1 := &num1
// 	fmt.Println(*addressTonum1) // it will give us the value that is stored in &num1
// 	fmt.Printf("Type of addressTonum1 %T\n", addressTonum1)
// 	fmt.Printf("Address of addressTonum1 %v\n", &addressTonum1)
// 	fmt.Println(*&addressTonum1)

// }

func main() {
	// Remember two operators while working with pointers
	// 1- & (amperson) - also called "address of"
	// 1- * (asterik) - also called "dereferencing"
	fmt.Println("Working and learning about pointers in golang!")
	// Regular var declaration and initialization using shorthand operator
	i, j := 42, 2701

	fmt.Println("Regular variables being printed out: ", i, j)

	// Print the address of variable i and j
	fmt.Printf("Address of i -> %v and j -> %v in memory allocation\n", &i, &j)

	// Now, if we assign "address of i" to "p" (new variable) like below:
	// It is now pointer
	p := &i

	// When we print out this p it will show the same value (address of i)
	fmt.Printf("Value of p -> %v\n", p)
	fmt.Printf("Address of p -> %v\n", &p)

	fmt.Println(*p)
	// In this case, the value of p is the address of i and...
	// *p (star p) is the value 'at' that address which is the value of i.
	// In Urdu:
	// p ki value -> address of i
	// *p ki value -> iski value wohi hogi jo is k address par assigned hai. is (*p) k address par i ki value assigned hai jo k 42 hai.

	// Question:
	// What happens if we try to change *p. For example:
	*p = 21
	// is case main "i" ki value 42 se change ho kar 21 ho jayegi.
	// Because *p ki value wo value hai jo us k address main assigned hai or wo value i ki value hai.
	fmt.Println("Value changed after assigning new value to *p")
	fmt.Println(i)

}
