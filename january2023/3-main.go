package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Todo App!")

	tasks := []string{}

	for true {

		fmt.Println("Add tasks: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		f_text := strings.TrimSpace(input)

		if f_text != "" {

		}

	}

}

func checkNilErr(err error) {
	panic(err)
}
