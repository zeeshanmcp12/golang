package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Saving task in a file")

	tasks := []string{}

	fmt.Print("Add Tasks: ")

	file, err := os.Create("./11-2-tasks.txt")
	checkNilErr(err)

	for i := 0; true; i++ {

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		f_txt := strings.TrimSpace(input)

		if f_txt != "done" {
			tasks = append(tasks, f_txt)
			fmt.Print("Type done or add task: ")

			continue
		} else if f_txt == "done" {
			fmt.Println("Thank you for adding tasks: ")
			for j := range tasks {
				// fmt.Printf("Task#%v, %v\n", j, val)
				io.WriteString(file, tasks[j])
			}
			break

		} else {
			fmt.Println("Unexpected error!")

		}
	}
	defer file.Close()
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
