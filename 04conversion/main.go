package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Please rate the course: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thank you for the rating:", input)
	fmt.Printf("Type of input var is %T", input)
	fmt.Println("\n------------------------------------")
	fmt.Println("Converstion started!")
	fmt.Println("------------------------------------")
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your ratings!", numRating+1)
	}

	fmt.Printf("\nType of input var is %T", numRating)

}
