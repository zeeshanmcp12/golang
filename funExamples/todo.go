package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Todo app")
	todos := []string{}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Add todo: ")
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	trimSpace := strings.TrimSpace(input)
	todos = append(todos, trimSpace)
	fmt.Println(todos)

}
