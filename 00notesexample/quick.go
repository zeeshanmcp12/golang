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

	var x, y int = 100, 9
	x /= y
	fmt.Println(x)
	x %= y
	fmt.Println(x)

}
