package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Conversion in golang!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter any number: ")

	input, _ := reader.ReadString('\n')
	fmt.Printf("Type of input is %T\n", input)

	conToInt, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Type of converted is %T and the value is %v", conToInt, conToInt)
	}
}
