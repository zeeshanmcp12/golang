package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// todo app logic
/*
func main() {
	fmt.Println("New File for practice!")
	fmt.Println()


		for i := 1; i < 10+1; i++ {
			fmt.Printf("%v x %v = %v\n", tableNum, i, tableNum*i)
		}

		fmt.Printf("Enter any number to make table: ")
		tableNum := userInput()

		for i := 1; i < 10+1; i++ {

			// anotherInt := i * 2
			fmt.Printf("%v x %v = ", tableNum, i)
			result := userInput()

			// Working logic
			if result == tableNum*i {
				fmt.Println("result executed")

			} else {
				fmt.Printf("Wrong Answer!\n%v x %v = %v\n", tableNum, i, userInput())
			}

		}

	todos := []string{}
	fmt.Printf("Add task: ")

	for i := 0; true; i++ {

		f_text := strings.TrimSpace(stringInput())

		// fmt.Printf("Add task or Enter done: ")
		if f_text != "done" {
			if f_text != "" {

				fmt.Printf("Add task or enter done: ")
				todos = append(todos, f_text)
				continue

			} else {
				fmt.Printf("Invalid input. Add task or enter done: ")

			}

		} else if f_text == "done" {
			fmt.Println()
			fmt.Println("Here is your tasks:")
			for j, val := range todos {
				fmt.Printf("%v %v\n", j+1, val)
			}
			break
		}
	}
}*/

// working with files
/*
func main() {

	fmt.Println("Working with Files in golang")
	content := "Magic! Magic! Magic!"

	file, err := os.Create("./practiceFile.txt")
	CheckNilErr(err)

	io.WriteString(file, content)

	defer file.Close()
	readFile("./practiceFile.txt")

}*/

const url = "https://acloudtechie.com"

func main() {
	fmt.Println("Handling web request in golang!")

	response, err := http.Get(url)
	CheckNilErr(err)
	fmt.Printf("Type of response is %T", response)
	fmt.Println()

	// fmt.Printf("Site status: %v", string(response.Status))
	defer response.Body.Close()

	res, err := ioutil.ReadAll(response.Body)
	fmt.Println(response.Header)

	content := res

	file, err := os.Create("./disclaimer.html")
	CheckNilErr(err)
	fmt.Println("File has been created and written with response.")

	io.WriteString(file, string(content))

}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)

	CheckNilErr(err)

	fmt.Printf("This data is retrieved from file:\n%v", string(data))

}

func userInput() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strToInt(input)

}

func strToInt(strText string) int {
	input, err := strconv.Atoi(strings.TrimSpace(strText))
	CheckNilErr(err)

	return input
}

func stringInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
