package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in golang!")

	var query string

	fmt.Print("Please ask for time: ")
	fmt.Scanf("%s\n", &query)

	// timeNow := time.Date(2006, time.January, 02, 00, 00, 00, 00, time.Local)

	fmt.Println("Time now is: ", time.Now())

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("2006-01-02 Monday"))

	createdDate := time.Date(2021, time.December, 03, 07, 00, 00, 00, time.Local)
	fmt.Println(createdDate)
	fmt.Println("----------------------------------------------")
	fmt.Println("CNIC Issue Date:", createdDate.Format("02-01-2006 15:04 Monday"))
	fmt.Println("----------------------------------------------")
}
