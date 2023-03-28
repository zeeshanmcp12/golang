package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Print execution time in golang!")
	vanquish()
}

func vanquish() {
	fmt.Println("vanquishing the Dragon")
	start := time.Now()

	time.Sleep(2 * time.Second)
	timeElapsed := time.Since(start)

	fmt.Printf("Dragon vanquished in %s\n", timeElapsed)

}
