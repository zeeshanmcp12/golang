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
	// var a, b = 100, 5
	// switch a {
	// case a/b == 10:
	// 	fmt.Println("10")
	// case a/b == 20:
	// 	fmt.Println("20")
	// case a/b == 10:
	// 	fmt.Println("30")
	// default:
	// 	fmt.Println("default")
	// }

	// for i := 0; i <= 5; i++ {
	// 	fmt.Println(i * i)
	// 	if i == 3 {
	// 		continue
	// 	}
	// }

	// var tableNum int
	// fmt.Printf("Enter table number: ")

	// fmt.Scanf("%d", &tableNum)

	// for i := 1; i < 10+1; i++ {
	// 	fmt.Printf("%v x %v = %v\n", tableNum, i, tableNum*i)
	// 	// if i == 4 {
	// 	// 	break
	// 	// }
	// }

	// days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// var leaveDay string

	// fmt.Printf("Apply for leave: ")
	// fmt.Scanf("%s", &leaveDay)
	// indexToRemove := days[:5]

	// days = append(indexToRemove)

	// for _, val := range days {
	// 	// fmt.Printf("Index: %v Day:%v\n", d, val)
	// 	fmt.Printf("Apply for leave: ")
	// 	fmt.Scanf("%s", &val)
	// 	if val == "Saturday" || val == "Sunday" {
	// 		fmt.Println("It's a weekend, apply for other days.")

	// 		break
	// 	} else if val != "Saturday" && val != "Sunday" {
	// 		// days = append(indexToRemove)
	// 		fmt.Printf("Remaining Days:%v\n", append(days[:5]))
	// 		// fmt.Printf("You applied for %v\n", leaveDay)
	// 		break

	// 	} else {
	// 		fmt.Println("Invalid input!")
	// 	}

	// }

	// arr := [10]int{10, 20, 30, 50}
	// fmt.Println(arr[0])
	// fmt.Println(arr[2])
	// fmt.Println(arr[4])
	// fmt.Println(arr[8])
	// fmt.Println(arr[10])

	// arr := [5]string{"a", "b", "c", "d", "e"}
	// slice := arr[:4]
	// fmt.Println(arr)
	// fmt.Println(slice)
	// slice[1] = "x"
	// fmt.Println(arr)
	// fmt.Println(slice)

	// Need to understand why this is the output
	// arr := [5]int{10, 20, 90, 70, 60}
	// slice := arr[:3]
	// fmt.Println(cap(slice))
	// new_slice := append(slice, 100, 200)
	// fmt.Println(cap(new_slice))

	// Need to understand why this is the output
	// arr := [5]int{10, 20, 90, 70, 60}
	// slice := arr[:3]
	// fmt.Println(cap(slice)) // 5

	// slice_2 := make([]int, 5, 20)
	// new_slice := append(slice, slice_2...)
	// fmt.Println(cap(new_slice)) // 10

	//It throws an error.
	// arr := [5]int{10, 20, 90, 70, 60}
	// slice := append(arr[:2], arr[3:])
	// fmt.Println(slice)

	// Need to understand why this is the output
	// arr := []int{10, 20, 90, 70, 60}
	// slice := make([]int, 10)
	// num := copy(slice, arr)
	// fmt.Println(slice) // [10 20 90 70 60 0 0 0 0 0]
	// fmt.Println(num)   // 5

	// Need to understand why this is the output
	// arr := []int{10, 20, 90, 70, 60}
	// slice := make([]int, 10)
	// copy(slice, arr)
	// slice[1] = 1000
	// fmt.Println(arr)   // [10 20 90 70 60]
	// fmt.Println(slice) // [10 1000 90 70 60 0 0 0 0 0]

	// arr := [10]int{10, 20}
	// // [10 1000 90 70 60 50 40 30 20 10]
	// slice := arr[2:8]       // 2 is 90 and 8 is 20 but not inclusive so this will be 30
	// fmt.Println(len(slice)) // 6 -> because from 90 to 30, the length will be 6
	// fmt.Println(cap(slice)) // 8 -> after 30 we still have 2 elements to it will be 6+2=8

	// ascii_codes := map[string]string{}
	// ascii_codes["A"] = 65
	// fmt.Println(ascii_codes) // Error: cannot use 65 (type untyped int) as type string in assignment

	// - create a map using make() function with key data type as string, and value data type as int.
	// - add the following key_value pairs to it
	// ("A", 65)
	// ("F", 70)
	// ("K", 75)
	// - delete the key "F"
	// - print the map

	// ------------------------------
	// var myMap = make(map[string]int)
	// myMap["A"] = 65
	// myMap["F"] = 70
	// myMap["K"] = 75

	// delete(myMap, "F")
	// fmt.Println(myMap) // output: map[A:65 K:75]
	// ------------------------------

	// ------------------------------
	// ascii_codes := make(map[string]int)
	// ascii_codes["A"] = 65
	// ascii_codes["F"] = 70
	// ascii_codes["K"] = 75
	// fmt.Println(ascii_codes)

	// ascii_codes = make(map[string]int)
	// ascii_codes["U"] = 85
	// fmt.Println(ascii_codes)
	// ------------------------------

	ascii_codes := make(map[string]int, 10)
	ascii_codes["A"] = 65
	ascii_codes["F"] = 70
	ascii_codes["K"] = 75
	fmt.Println(len(ascii_codes))
	ascii_codes = make(map[string]int)
	ascii_codes["U"] = 85
	fmt.Println(len(ascii_codes))
}
