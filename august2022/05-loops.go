package main

import (
	"fmt"
)

func main() {
	fmt.Println("Using loops with slice")

	// days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Printf("%v \n", days[i])
	// }

	// for _, value := range days {
	// 	fmt.Printf("%v \n", value)
	// }

	// in this case it will return the value itself and not index
	// for i := range days {
	// 	// fmt.Println(days[i])
	// 	fmt.Printf("days: %v\n", days[i])
	// }

	// in this case, it will print the value and skip index because we used blank identifier
	// for _, value := range days {
	// 	fmt.Printf("Value is %v\n", value)
	// }

	var (
		yourNum int
	)
	fmt.Print("Enter any number: ")
	fmt.Scanf("%d", &yourNum)
	fmt.Println("Your number: ", yourNum)

	for yourNum < 10+1 {
		// in this case when number is matched to 5 it will break the loop
		// if yourNum == 5 {
		// 	// fmt.Println("Congratulations! Number Matched")
		// 	break
		// }

		// in this case, if condition matches then first we will upgrade the value and then the number will be skipped
		// if yourNum == 5 {
		// 	yourNum++
		// 	continue
		// }

		// in this case the loop will be terminated
		// if yourNum == 5 {
		// 	yourNum++
		// 	break
		// }

		if yourNum == 8 {
			goto webLink
		}

		fmt.Println("Number is: ", yourNum)
		yourNum++
	}

webLink:
	fmt.Println("Jumping to acloudtechie.com")
}
