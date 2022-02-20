package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Practice..Pactice..Practice!")

	// ------------- Print table of any number by taking input ----------------
	/*
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter any number to calulate it's table: ")
		input, _ := reader.ReadString('\n')
		tableNum, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Thank you for entering %v", tableNum)

			for i := 0; i < 10+1; i++ {
				fmt.Printf("%v x %v = %v\n", tableNum, i, tableNum*i)
			}

		}*/
	// ------------- Print table of any number by taking input ----------------

	// ------------- Daily Quotes ----------------
	birthdayQuotes := []string{"Happy Birthday\n", "Cuties born in Feb!\n", "Enjoy your day"}
	bestWishes := []string{"All the best!\n", "Don't give up\n", "Alway try to not getting fear of"}

	for i := 0; true; i++ {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Select any number: [1]Birthday Quotes, [2]Best Wishes: ")
		input, _ := reader.ReadString('\n')
		c_quotes, _ := strconv.Atoi(strings.TrimSpace(input))
		if c_quotes == 1 {
			fmt.Printf("Birthday Quotes:\n%v\n", birthdayQuotes)
			fmt.Println()
			continue

		} else if c_quotes == 2 {
			fmt.Printf("Best Wishes:\n%v\n", bestWishes)
			break

		} else {
			fmt.Println("Else statement executed!")
		}
	}
}
