package main

import "fmt"

func main() {
	fmt.Println("Array in golang!")

	var grades = [4]int{20, 30, 40, 50}
	fmt.Println(grades)
	fmt.Println("Length of array is:", len(grades))

	grades[1] = 35
	fmt.Println("Grade has been changed:", grades)

	// for i := 0; i < len(grades); i++ {

	// 	fmt.Printf("array element: %v\n", grades[i])

	// }

	var arr = [3][2]int{{10, 30}, {5, 20}, {15, 66}}
	fmt.Println(arr[0][1])
	fmt.Println(len(arr))

	for i := 0; i < len(arr); i++ {
		fmt.Println(i, "=>", arr[i][0])
		fmt.Println(i, "=>", arr[i][1])
		// fmt.Println(i, "=>", arr[i][1])
	}

}
