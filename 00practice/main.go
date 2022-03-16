package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
func main() {
	fmt.Println("Reading notes to rephrase the learning!")
	fmt.Println("Taking input from user!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter any text: ")
	input, _ := reader.ReadString('\n')
	converted, _ := strconv.Atoi(strings.TrimSpace(input))
	fmt.Printf("Type of string is %T and value is %v", converted, input)

}*/

func main() {
	fmt.Println("Another main function!")
	rand.Seed(time.Now().UnixNano())
	randNumber := rand.Intn(5)
	fmt.Println(randNumber)

	currentTime := time.Now()
	fmt.Println(currentTime.Format("Monday, 2-1-2006"))

	currentDate := time.Date(2006, time.March, 2, 00, 00, 00, 00, time.Local)
	fmt.Println(currentDate)
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
