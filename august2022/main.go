package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Building applications in golang!")

	tasks := []string{}

	fmt.Printf("Add tasks: ")

	for i := 0; true; i++ {

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		f_txt := strings.TrimSpace(input)

		if f_txt != 'done' {
			if f_txt != "" {
				
			}
			
		}

	}
}
