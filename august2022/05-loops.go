package main

import "fmt"

func main() {
	fmt.Println("Using loops with slice")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Printf("%v \n", days[i])
	// }

	// for _, value := range days {
	// 	fmt.Printf("%v \n", value)
	// }

	// in this case it will return the value itself and not index
	for i := range days {
		// fmt.Println(days[i])
		fmt.Printf("days: %v\n", days[i])
	}

	// in this case, it will print the value and skip index because we used blank identifier
	for _, value := range days {
		fmt.Printf("Value is %v\n", value)
	}

}
