package main

import "fmt"

func main() {
	fmt.Println("Array in golang!")

	var weekendDays1 string
	var weekendDays2 string

	var weekdays [7]string
	weekdays[0] = "Monday"
	weekdays[1] = "Tuesday"
	weekdays[2] = "Wednesday"
	weekdays[3] = "Thursday"
	weekdays[4] = "Friday"
	// weekdays[0] = "Monday"

	fmt.Println("Total elements in an array: ", len(weekdays))
	fmt.Println("Days in a week which we work for: ", weekdays)

	var weekend = [2]string{"Saturday, Sunday"}
	fmt.Println("Elements in weekend array: ", len(weekend))
	fmt.Println("We are off on: ", weekend)

	fmt.Printf("What days you take off, day1: ")
	fmt.Scanf("%s\n", &weekendDays1)

	fmt.Printf("What days you take off, day2: ")
	fmt.Scanf("%s\n", &weekendDays2)

	if weekendDays1 == "Saturday" && weekendDays2 == "Sunday" {
		fmt.Println("Happy Weekend!")

	} else {
		fmt.Println("These are alternate days, so Enjoy your off days!")
	}

}
