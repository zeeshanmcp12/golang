package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time formatting in golang!")

	// currentTime := time.Now()
	// fmt.Println("Current time is: ", currentTime)

	// formttedTime := currentTime.Format("01-02-2006")
	// fmt.Println(formttedTime)

	createdTimeDate := time.Date(1990, time.February, 15, 00, 00, 00, 00, time.Local)
	fmt.Println(createdTimeDate)
	formatted := createdTimeDate.Format("Monday, 01 Feb, 2006")
	fmt.Println(formatted)

	var year int
	var month int
	var date int

	fmt.Printf("Enter year of your birth: ")
	fmt.Scanf("%v\n", &year)

	fmt.Printf("Enter Month of your birth: ")
	fmt.Scanf("%v\n", &month)

	fmt.Printf("Enter date of your birth: ")
	fmt.Scanf("%v\n", &date)

	completeDate := time.Date(year, time.Month(month), date, 00, 00, 00, 00, time.Local)

	fmt.Println("Thank you for entring:", completeDate.Format("02 Jan, 2006, Monday"))
}
