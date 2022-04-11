package main

import "fmt"

func main() {
	// var grades int = 42
	// fmt.Printf("variable grades = %v is of type %v \n", grades, grades)

	// var grades int = 42
	// fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
	// const name = "Harry"
	// fmt.Print(name)

	// var name = "Harry"
	// fmt.Print(name)

	// var name string = "Harry"
	// fmt.Print(name)

	// var a int = 12
	// var b int = 12
	// fmt.Println(a <= b)

	// a = 20
	// fmt.Println(a <= b)

	// b = 100
	// fmt.Println(a <= b)

	// c := 0
	// fmt.Println(a <= c)

	// var a, b bool = true, false
	// fmt.Println(a && b)
	// fmt.Println(a || b)

	// var a, b bool = false, false
	// fmt.Println(a && b)
	// fmt.Println(a || b)

	// var a, b bool = false, true
	// fmt.Println(!a)
	// fmt.Println(b)

	// var a bool = false
	// result := 10 > 50
	// fmt.Println(!(a && result))

	// var a bool = true
	// result := 10 > 50
	// fmt.Println(!(a || result))

	// var a, b bool = false, false
	// fmt.Println(a && b)
	// fmt.Println(a || b)

	// var x, y int = 100, 9
	// x /= y
	// fmt.Println(x)
	// x %= y
	// fmt.Println(x)

	// bitwise AND and OR
	// var x, y int = 100, 90
	// fmt.Println(x & y)
	// fmt.Println(x | y)

	// need to validate
	// var x, y int = 100, 90
	// fmt.Println(!(((x + y) >> 2) == 47))

	// Tricky question in kodekloud lab.
	// If below code is not in editor then we need to find out either there will be an error or any result.
	// This will through an error because switch contains "int" value while cases contains bool and it is not allowed to convert untyped bool to int. That's why it will give error.
	var a, b = 100, 5
	switch a {
	case a/b == 10:
		fmt.Println("10")
	case a/b == 20:
		fmt.Println("20")
	case a/b == 10:
		fmt.Println("30")
	default:
		fmt.Println("default")
	}

}
