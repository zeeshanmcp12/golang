package main

import "fmt"

// Without using / implementing interface
type Book struct {
	Title string
	Price int
}

type Drink struct {
	Name  string
	Price float64
}

func (b Book) printInfo() {
	fmt.Printf("Title -> %s\nPrice -> %d\n", b.Title, b.Price)
}

func (d Drink) printInfo() {
	fmt.Printf("Name -> %s\nPrice -> %.2f\n", d.Name, d.Price)
}

func main() {
	fmt.Println("Interfaces in golang!")

	book := Book{
		Title: "The Game Changer",
		Price: 1000,
	}

	drink := Drink{
		Name:  "Mint Margritta",
		Price: 100,
	}

	book.printInfo()
	drink.printInfo()
}
