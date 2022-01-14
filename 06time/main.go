package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("2006-01-02 Monday"))

	createdDate := time.Date(2021, time.December, 03, 07, 00, 00, 00, time.Local)
	fmt.Println(createdDate)
	fmt.Println("----------------------------------------------")
	fmt.Println("CNIC Issue Date:", createdDate.Format("02-01-2006 15:04 Monday"))
	fmt.Println("----------------------------------------------")

}
