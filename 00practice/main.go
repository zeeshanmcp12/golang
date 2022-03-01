package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Practice about functions in golang")

	// writing function inside function
	output := func() {
		fmt.Println("function call from inside of function")
	}
	output()

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	f_text, err := strconv.Atoi(strings.TrimSpace(input))

	CheckNilErr(err)

	fmt.Println(f_text)

}
