package main

import "fmt"

type Book struct {
	Title string
	Price int
}

type Drink struct {
	Name  string
	Price float64
}
type Chips struct {
	Name  string
	Price float64
}

type Printer interface {
	returnInfo() string
	// returnthirdInfo() string
}

func (b Book) returnInfo() string {
	// fmt.Printf("Title -> %s\nPrice -> %d\n", b.Title, b.Price)
	return "returning book info"
}

func (b Drink) returnInfo() string {
	return "returning drink info"
}

func (c Chips) returnInfo() string {
	return "returning chips info"
}

func printData(p Printer) {

	fmt.Println(p.returnInfo())
	// fmt.Println(p.returnthirdInfo())

}

func main() {
	fmt.Println("Interfaces in golang!")

	book := Book{
		Title: "The Game Changer",
		Price: 1000,
	}
	printData(book)

	drink := Drink{
		Name:  "Mint Margritta",
		Price: 100,
	}
	printData(drink)

	chips := Chips{
		Name:  "Cheese Fries",
		Price: 200,
	}
	printData(chips)

	// book.printInfo()
	// drink.printInfo()
}
