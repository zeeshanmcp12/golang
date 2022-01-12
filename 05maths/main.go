package main

import (
	// "math/rand"
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// Math in golang
	// var firstNumber int = 6
	// var secondNumber int = 5
	// fmt.Println(firstNumber + secondNumber)

	// Take Input from user
	/* 	reader := bufio.NewReader(os.Stdin)
	   	fmt.Printf("Enter any number: ") // Make sure to not enter string.
	   	input, _ := reader.ReadString('\n')
	   	fmt.Printf("Type of input var is %T \n", input)

	   	// Convert string (from input var) into Int
	   	convertToInt, err := strconv.Atoi(strings.TrimSpace(input))

	   	if err != nil {
	   		fmt.Println(err)
	   	} else {
	   		// rand.Seed(time.Now().UnixNano()) // Incase you want to generate random number without giving seed below.

	   		// random number
	   		rand.Seed(int64(convertToInt))
	   		fmt.Println("Random number: ", rand.Intn(11))
	   		fmt.Printf("Type of convertToInt var is %T", convertToInt)
	   	} */

	/*
		// Generating random numbers using math/rand
		rand.Seed(time.Now().UnixNano()) // since time is keep on changing so this time Seed will surely generate random numbers.
		fmt.Println(rand.Intn(6)) */

	// Generating random numbers using crypto/rand
	randomNumByCrypto, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(randomNumByCrypto)
}
