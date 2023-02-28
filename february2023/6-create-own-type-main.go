package main

import "fmt"

var num1 int

type ownType int

var myType ownType

func main() {
	fmt.Println("Create your own type in golang!")
	fmt.Println("We can create our own type in golang.")

	num1 = 42
	fmt.Println(num1)
	fmt.Printf("Type of num1 -> %T\n", num1)

	myType = 52
	fmt.Println(myType)
	fmt.Printf("Type of myType -> %T\n", myType)

	// We cannot do following thing:
	// num1 = myType
	// This will give error because golang is static typed so we cannot myType (variable of type ownType) as int value in assignment

	// We need to convert myType to "int"
	num1 = int(myType)
	fmt.Printf("Type of num1 after conversion -> %T\n", num1)

}
