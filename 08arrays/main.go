package main

import "fmt"

func main() {
	// fmt.Println("Learning about arrays in golang!")
	// var fruitList [4]string
	// fruitList[0] = "Apple"
	// fruitList[1] = "Peach"
	// fruitList[3] = "Mango"
	// fmt.Println("FruitList is: ", fruitList)
	// fmt.Println("Length of array is: ", len(fruitList))

	// var vegList = [3]string{"Potatoe", "Mushrom", "Tomato"}
	// fmt.Println("Vegetable List is: ", vegList)
	// fmt.Println("Vegetable List is: ", len(vegList))
	var todos = [4]string{"Recite Quran", "Workout", "Sprint Task"}
	fmt.Println("My todos are: ", todos[0])
	fmt.Println("My todos are: ", todos[3])
	fmt.Println(len(todos[0]))

}
