package main

import "fmt"

func main() {
	fmt.Println("Array and Slice in golang!")

	var fruitList [5]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Banana"
	fruitList[3] = "Guava"
	fruitList[4] = "Peach"

	fmt.Println(fruitList)

	for i, val := range fruitList {
		fmt.Printf("%v -> %v\n", i, val)
	}

	user := User{"Abdullah", 32, "abd@go.dev"}

	fmt.Println(user.Email)
	fmt.Println("----- Email updated ------")
	user.updateEmail()

}

type User struct {
	Name  string
	Age   int
	Email string
}

func (u *User) updateEmail() {

	u.Email = "xyz@go.dev"
	fmt.Println(u.Email)

}
