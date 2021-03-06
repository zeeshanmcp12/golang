package main

import (
	"fmt"
)

func main() {
	fmt.Println("Learning about Slices in golang!")
	// ==================================================================================
	/*	var sliceList = []string{"Apple", "Mango", "Banana"}
		fmt.Printf("Type of sliceList var is %T\n", sliceList)
		fmt.Printf("--------------------------------\n")
		fmt.Printf("sliceList before adding values\n")
		fmt.Println(sliceList)
		fmt.Printf("--------------------------------\n")
		fmt.Printf("sliceList after adding values\n")
		sliceList = append(sliceList, "Peach", "Orange")
		fmt.Println(sliceList)
		fmt.Printf("--------------------------------\n")
		fmt.Println("-----Slicing------")
		// fmt.Println("-----Slicing from 1 to not ending list------")
		// sliceList = append(sliceList[1:]) // Start from 1 which is Mango because Apple is on 0th position
		// fmt.Println(sliceList)
		// fmt.Println()
		// fmt.Println("-----Slicing from 1 to 3 elements------")
		// sliceList = append(sliceList[1:3]) // Start from 1 which is Mango, second element is Banana and third is Peach which is last one and non-inclusive.
		// fmt.Println(sliceList)
		fmt.Println("-----Slicing from 3 to whatever before it------")
		sliceList = append(sliceList[:3]) // start from the default value which is 0 (it's value is Apple), value on 1st place is Mango, 2nd place is Banana, 3rd place is Peach which is last range so non-inclusive.
		fmt.Println(sliceList)
		// ==================================================================================
	*/
	highScores := make([]int, 4)

	highScores[0] = 123
	highScores[1] = 234
	highScores[2] = 456
	highScores[3] = 678
	// fmt.Println(highScores)
	// highScores[4] = 891 // This will throw an error because we just defined 4 elements in slice.
	// fmt.Println(highScores)
	// let's try to use append method and add some more elements in slice.
	// highScores = append(highScores, 876, 654, 432, 321)
	// fmt.Println(highScores)

	// in above case, slice should not allow to add more values because we had defined 4 initially. But it allowed us and added more values into it.
	// The reason behind this is "reallocating the memory" in golang. This time, it realloced the memory and accumodated more values in slice.
	// When we use append method, it reallocates the memory that's why it added more values into it.
	// Initially we did great with memory allocation and defined only 4 elements but when append came, memory reallocation happend again. this comes up with lots of performance optimization benefits etc

	// sort.Ints(highScores) // sort is a package and Ints is a method of it. It sorts the value in slice.
	// sort is only available in slices so not in array.
	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores)) // This function returns if values in slice are sorted or not so it returns true or false.
	var cources = []string{"javascript", "Java", "C#", "golang", "swift"}
	fmt.Println("------Before removing item from slice-------")
	fmt.Println(cources)
	fmt.Println()
	var index int = 2
	cources = append(cources[:index], cources[index+1:]...)
	fmt.Println("------After removing item from slice-------")
	fmt.Println(cources)
	fmt.Println()

}
