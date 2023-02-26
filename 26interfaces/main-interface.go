package main

import "fmt"

type Animal interface {
	printData() string
}

type Cat struct {
	Name string
}

func (c Cat) printData() string {
	return c.Name
}

type Dog struct {
	Name string
}

func (d Dog) printData() string {
	// return "Nothing!"
	return d.Name
}

func main() {
	fmt.Println("Interface in golang!")

	var myAnimal Animal
	myAnimal = Cat{Name: "Ruby"}
	fmt.Println(myAnimal.printData())

	myAnimal = Dog{Name: "Puppy"}
	fmt.Println(myAnimal.printData())

}
