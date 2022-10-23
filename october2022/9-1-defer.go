package main

import "fmt"

func main() {
	fmt.Println("More examples of defer in golang!")
	fmt.Println("-------------------")
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")
	fmt.Println("Hello")
	tryDefer()
	fmt.Println()

}

/*
stack -> Last in
1- World
2- One
3- Two
4- Three

stack -> Fist out -> when it gets printed out
Hello -> statement without defer
tryDefer -> function without defer keyword -> 43210
1- Three
2- Two
3- One
4- World
*/

func tryDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%v", i)

	}
}
