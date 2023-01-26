package main

import (
	"fmt"
	"strings"
)

func main() {
	greeterHello := "hello"
	fmt.Println(greeterHello)

	modifyGreeter := &greeterHello
	fmt.Printf("Value of greeterHello -> %v\n", greeterHello)
	*modifyGreeter = "world"
	fmt.Printf("Value of greeterHello -> %v \n", greeterHello)
	*modifyGreeter = "world"
	*modifyGreeter = strings.ToUpper(greeterHello)
	fmt.Printf("Value of greeterHello -> %v \n", greeterHello)
	fmt.Println("----- Passing by Reference to function -----")
	modifyGreeterFunc(&greeterHello)
	fmt.Println("Value of greeterHello after using `Passing by Reference`:", greeterHello)

}

func modifyGreeterFunc(str *string) {
	*str = "Hello World"
}
