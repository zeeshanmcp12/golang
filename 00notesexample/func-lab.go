// package main

// import "fmt"

// func printString(str string) {
// 	fmt.Printf("%q ", str)
// }

// func printInt(i int) {
// 	fmt.Printf("%d ", i)
// }

// func printFloat(f float64) {
// 	fmt.Printf("%.2f ", f)
// }
// func main() {
// 	printString("browser")
// 	defer printInt(32)
// 	defer printFloat(0.24)
// 	printString("chrome")
// 	printInt(90)
// 	defer printFloat(89)
// 	printInt(900)
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func getString(str string) (string, string) {
// 	return strings.ToLower(str), strings.ToUpper(str)
// }

// func main() {
// 	_, lower := getString("BROWSER")
// 	fmt.Println(lower)
// }

package main

import (
	"fmt"
)

func main() {
	fmt.Println(f1())
}

func f1() int {
	return f2()
}

func f2() int {
	return 1
}
