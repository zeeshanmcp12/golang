package main

import "fmt"

func main() {
	fmt.Println("Defer in golang!")
	fmt.Println("-----------------------")
	defer fmt.Println("With defer keyword")
	defer fmt.Println("One With defer keyword")
	fmt.Println("Without defer keyword")
}

// "With defer keyword" -> this gets into stack first -> after that below will get into stack
// "One With defer keyword" -> last in > means it gets into stack at very last. So it will print on top after regular statement. for example:
// "Without defer keyword" -> "One With defer keyword" -> "With defer keyword"
// This also means last in first out (LIFO) -> because this ("One With defer keyword") statement came into line at very last but printed out at very top (after regular statement)

/*
Output of above program:

Defer in golang!
-----------------------
Without defer keyword
One With defer keyword
With defer keyword
*/
