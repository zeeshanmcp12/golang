package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Practice..Pactice..Practice!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter any number to calulate it's table: ")
	input, _ := reader.ReadString('\n')
	tableNum, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Thank you for entering %v", tableNum)

		for i := 0; i < 10+1; i++ {
			fmt.Printf("%v x %v = %v\n", tableNum, i, tableNum*i)
		}

	}
}
