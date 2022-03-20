package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
func main() {
	fmt.Println("Practice from 6 to 10")
	fmt.Printf("Enter any number: ")
	anything := userInputtoInt()
	fmt.Printf("Type is %T: ", anything)
	fmt.Println(anything)

	fmt.Printf("Enter any text: ")
	anyText := onlyString()
	stringText := strings.TrimSpace(anyText)
	fmt.Printf("Type is %T: ", stringText)

}*/
/*
func main() {
	fmt.Println("Practice 6 to 10")
	fmt.Println("-----------------------")
	fmt.Printf("Enter date of your birth: ")
	date := userInputtoInt()

	fmt.Printf("Enter month of your birth: ")
	month := userInputtoInt()

	fmt.Printf("Enter year of your birth: ")
	year := userInputtoInt()

	birthDate := time.Date(year, time.Month(month), date, 00, 00, 00, 00, time.Local)
	fmt.Printf("Happy Birthday on %v\n", birthDate.Format("02 Jan, 2006"))
	fmt.Printf("It was %v!\n", birthDate.Format("Monday"))

}*/

func main() {
	fmt.Println("Practice 6 to 10")
	/*var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Banana"
	fruitList[3] = "Peach"

	fmt.Println("Total elements in array are: ", len(fruitList))
	fmt.Println("Fruilist:", fruitList)
	fmt.Printf("Total characters in %v are %v\n", fruitList[1], len(fruitList[1]))
	*/

	// for i, v := range fruitList {
	// 	fmt.Printf("Index is %v and value is %v\n", i, v)

	// }

	// var todos = [4]string{"Recite Quran,", "Learn to code,", "Write some words"}
	// fmt.Println(todos)

	// array := [...]int{6, 4, 3, 5, 4, 34, 8, 9, 20}
	// // fmt.Println(len(array))
	// for i := 0; i < len(array); i++ {
	// 	fmt.Printf("index is %v, value is %v\n", i, array[i])

	// }
	/*
		tableValue := 2

		for i := 1; i < 10+1; i++ {
			fmt.Printf("%v x %v = %v\n", tableValue, i, tableValue*i)

		}*/

	//slice in golang
	/*
		var sliceList = []string{".NET", "Java", "golang"}
		fmt.Printf("Type of sliceList is %T\n", sliceList)

		sliceList = append(sliceList, "Ruby", "JavaScript")
		fmt.Println(sliceList)

		sliceList = append(sliceList[1:3])
		fmt.Println(sliceList)*/

	/*
		var intsForSort = []int{}
		intsForSort = append(intsForSort, 3, 2, 4, 6, 3, 6, 5, 9, 10, 8)
		sort.Ints(intsForSort)
		fmt.Println(intsForSort)*/

	//slice with make() function
	var sliceWithMake = make([]int, 4)
	sliceWithMake[0] = 120
	sliceWithMake[1] = 220
	sliceWithMake[2] = 320
	sliceWithMake[3] = 420
	// fmt.Println(sliceWithMake)

	// sliceWithMake = append(sliceWithMake, 520, 620, 720)
	// fmt.Println(sliceWithMake)

	// map
	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages[".Net"] = "dotnet"
	languages["PS"] = "PowerShell"

	fmt.Println("JS stands for", languages["JS"])

}

func userInputtoInt() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strToInt(input)
}

func strToInt(strText string) int {
	converted, err := strconv.Atoi(strings.TrimSpace(strText))
	CheckNilErr(err)
	return converted
}

func onlyString() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return input

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}

}
