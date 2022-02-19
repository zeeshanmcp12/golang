package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Print Table")
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter any number to print table: ")
	input, _ := reader.ReadString('\n')
	c_table_num, err := strconv.Atoi(strings.TrimSpace(input))

	// table_num := c_table_num * 2

	if err != nil {
		fmt.Println(err)
	} else {

		for d := 1; d < 10+1; d++ {
			fmt.Printf("%v x %v = %v\n", c_table_num, d, c_table_num*d)
		}
	}

	fmt.Printf("Press enter to continue....")
	pressAnyKey, _ := reader.ReadString('\n')
	fmt.Println(pressAnyKey)
}
