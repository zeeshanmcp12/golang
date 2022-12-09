package main

import "fmt"

// func main() {
// 	fmt.Println("Hello World")

// 	var name string = "Sakina"

// 	for i, j := range name {
// 		fmt.Printf("%v%c\n", i, j)
// 	}
// }

// func main() {
// 	var number int
// 	fmt.Println("Enter any number: ")
// 	fmt.Scanf("%v", &number)

// 	for i := 1; i < 10+1; i++ {
// 		fmt.Printf("%v x %v = %v\n", number, i, number*i)
// 	}
// }

// func main() {
// 	fmt.Println("Enter any text: ")

// 	reader := bufio.NewReader(os.Stdin)

// 	input, _ := reader.ReadString('\n')
// 	fmt.Println("This is what you entered:", input)

// 	fmt.Printf("Type of input is %T\n", input)

// 	strtoInt, err := strconv.Atoi(strings.TrimSpace(input))

// 	fmt.Printf("Now, type of input is %T", strtoInt)

// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	fmt.Println("Array in golang!")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Banana"
	fruitList[3] = "Orange"

	fmt.Println("Elements in an array: ", fruitList)
	fmt.Println("Elements in an array: ", fruitList[1])

	for i := range fruitList {
		fmt.Printf("%v%v\n", i, fruitList[i])
	}
}
