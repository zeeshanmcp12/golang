package main

import "fmt"

func main() {
	fmt.Println("Array in golang!")

	// var grades = [4]int{20, 30, 40, 50}
	// fmt.Println(grades)
	// fmt.Println("Length of array is:", len(grades))

	// grades[1] = 35
	// fmt.Println("Grade has been changed:", grades)

	// for i := 0; i < len(grades); i++ {

	// 	fmt.Printf("array element: %v\n", grades[i])

	// }

	// multidimensional array
	/*var arr = [3][2]int{{10, 30}, {5, 20}, {15, 66}}
	fmt.Println(arr[0][1])
	fmt.Println(len(arr))

	for i := 0; i < len(arr); i++ {
		fmt.Println(i, "=>", arr[i][0])
		fmt.Println(i, "=>", arr[i][1])
		// fmt.Println(i, "=>", arr[i][1])
	}*/
	// var table = [4]int{2, 3, 4, 5}

	// calculateTable([]int{2, 3, 4})
	calTable([]int{2})
	// i := 0
	// for i < len(table) {
	// 	for d := 0; d < 10+1; d++ {
	// 		if !(table[d]*d == 10) {
	// 			fmt.Printf("%v x %v = %v\n", table[d], d, table[d]*d)
	// 			continue
	// 		}
	// 	}
	// 	i++
	// 	// if len(table) == 0 {
	// 	// 	break
	// 	// }
	// }

	// for index, val := range table {
	// 	if index != 10+1 {
	// 		fmt.Printf("%v x %v = %v\n", val, index, val*index)

	// 		continue

	// 	}

	// }

	// table := 2

	// for i := 1; i < 10+1; i++ {

	// 	fmt.Printf("%v x %v = %v\n", table, i, table*i)
	// }

}

func calculateTable(arr []int) int {
	res := 0
	// for d := 0; d < len(arr); d++ {
	for i := 0; i < 10+1; i++ {
		// res *= arr[i]
		res = arr[i] * i
		fmt.Printf("%v x %v = %v\n", res, i, res)
		// if i != 10 {
		// 	continue
		// }
	}
	// }
	return res
}

func calTable(arr []int) {
	for i := 0; i < 10+1; i++ {
		fmt.Printf("%v x %v = %v\n", arr[i], i, arr[i]*i)

	}
}
