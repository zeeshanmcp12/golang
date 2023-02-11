package main

import "fmt"

type Movie struct {
	name   string
	rating float64
}

func getMovie(n string, r float64) (m Movie) {
	m = Movie{
		name:   n,
		rating: r,
	}
	return
}

func increaseRating(m *Movie) {
	(*m).rating += 1.0

}

func main() {
	fmt.Println("sturct and pass struct by reference")
	mov := getMovie("nothing", 4.0)
	// fmt.Printf("Movie data is %+v", mov)

	fmt.Printf("%+v\n", mov)
	increaseRating(&mov)
	fmt.Printf("%+v", &mov)

}
