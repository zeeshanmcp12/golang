package main

import "fmt"

type Printer interface {
	printData() string
}

type Book struct {
	Title string
	Price int
}

type Drink struct {
	Name  string
	Price int
}

func (b Book) printData() string {
	return "Returning Book info"
}

func (d Drink) printData() string {
	return "Returning Drink info"
}

func main() {
	fmt.Println("Desribe interface in detail!")

	fmt.Println("An interface is a collection of method signatures.")

	var p Printer

	p = Book{}
	d := Drink{"Mint", 100}

	fmt.Println(p.printData())
	fmt.Println(d.printData())

}
