package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("---------------Something cool inside this Words library---------------")
	var birthday = []string{"Happy Birthday\n", "Enjoy your special day\n", "Cuties born in Feb"}
	var bestWishes = []string{"Wish you all the best\n", "May you bless with success\n", "Don't give up"}

	// fmt.Printf("Birthday:\n%v\n", birthday)
	// fmt.Printf("Best Wishes:\n%v", bestWishes)

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Select anyone to see something cool:\n[1]Birthday, [2]BestWishes: ")
	selection, _ := reader.ReadString('\n')

	convertToInt, _ := strconv.Atoi(strings.TrimSpace(selection))

	if convertToInt == 1 {
		// fmt.Println(convertToInt)
		fmt.Printf("\n==============Birthday Qoutes===============\n")
		fmt.Println(birthday)

	} else if convertToInt == 2 {

		fmt.Printf("\n==============Best Wishes Qoutes===============\n")
		fmt.Println(bestWishes)

	} else {

		fmt.Println("Enter correct value")
	}

}
