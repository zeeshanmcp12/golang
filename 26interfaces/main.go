package main

import "fmt"

type Animal interface {
	Speak() string
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func main() {
	fmt.Println("Interface and it's implementation in golang!")

	// Here we will implement the interface
	// We define the interface variable of type Animal (which is an interface)
	var myAnimal Animal

	// We create an instance of Cat and assign it to myAnimal (interface variable)
	myAnimal = Cat{Name: "Ruby"}

	//Finally, we call the Speak() method on myAnimal, which returns the expected string based on the type of animal it contains at the time of the call.
	fmt.Println(myAnimal.Speak())

	// We create an instance of Dog and assign it to myAnimal (interface variable)
	myAnimal = Dog{Name: "Puppy"}

	//Finally, we call the Speak() method on myAnimal, which returns the expected string based on the type of animal it contains at the time of the call.
	fmt.Println(myAnimal.Speak())

	/*
		Note that we don't need to know the concrete type of myAnimal at compile time in order to call its Speak() method. This is because we're using the Animal interface, which provides a common way to interact with both Dog and Cat objects regardless of their underlying type. This is a key feature of interfaces and allows for more flexible and extensible code.
	*/

}
