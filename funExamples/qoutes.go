package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("---------Best Qoutes---------")

	var birthday = []string{"Happy Birthday\n", "Enjoy your special day\n", "Cuties born in Feb"}
	var bestWishes = []string{"Wish you all the best\n", "May you bless with success\n", "Don't give up"}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Select anyone to see something cool:\n[1]Birthday, [2]BestWishes: ")
	selection, _ := reader.ReadString('\n')

	convertToInt, _ := strconv.Atoi(strings.TrimSpace(selection))

	switch convertToInt {
	case 1:
		fmt.Printf("\n==============Birthday Qoutes===============\n")
		fmt.Println(birthday)
	case 2:
		fmt.Printf("\n==============Best Wishes Qoutes===============\n")
		fmt.Println(bestWishes)
	default:
		fmt.Println("Enter correct value")

	}
}
