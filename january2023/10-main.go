package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Trim Prefix")

	strText := "Y2Mate.is - How computers work. Number systems-CkFtG6qQDD8-1080p-1655251585172.mp4"

	stringTrimmed := strings.TrimPrefix(strText, "Y2Mate.is - ")

	fmt.Printf("Type of stringTrimmed %T\n", stringTrimmed)
	fmt.Println("Type of stringTrimmed: ", stringTrimmed)
}
