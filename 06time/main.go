package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// presentTime := time.Now()
	// fmt.Println(presentTime)
	// fmt.Println(presentTime.Format("2006-01-02 Monday"))

	// createdDate := time.Date(2021, time.December, 03, 07, 00, 00, 00, time.Local)
	// fmt.Println(createdDate)
	// fmt.Println("----------------------------------------------")
	// fmt.Println("CNIC Issue Date:", createdDate.Format("02-01-2006 15:04 Monday"))
	// fmt.Println("----------------------------------------------")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter Year of your Birth: ")
	year, _ := reader.ReadString('\n')
	fmt.Printf("Enter Month of your Birth (only number): ")
	month, _ := reader.ReadString('\n')
	fmt.Printf("Enter Date of your Birth: ")
	day, _ := reader.ReadString('\n')

	convertYear, _ := strconv.Atoi(strings.TrimSpace(year))
	convertMonth, _ := strconv.Atoi(strings.TrimSpace(month))
	convertDay, _ := strconv.Atoi(strings.TrimSpace(day))

	createdAt := time.Date(convertYear, time.Month(convertMonth), convertDay, 00, 00, 00, 00, time.Local)

	// fmt.Println(createdAt)
	fmt.Println()
	fmt.Println("-------------------- Date of Birth --------------------")
	// fmt.Println()
	fmt.Println(createdAt.Format("02 Jan 2006"))
	fmt.Println()
	fmt.Println("-------------- Day of your Birth ------------")
	fmt.Println(createdAt.Format("It was: Monday. So, Happy Birthday on that day!"))
	fmt.Println()
	fmt.Printf("Press 'Enter' key to continue!")
	pressedKey, _ := reader.ReadString('\n')
	fmt.Println(pressedKey)
}
