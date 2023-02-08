package main

import "fmt"

type Movie struct {
	name   string
	rating float32
}

func getMovie(s string, r float32) (m Movie) {
	m = Movie{
		name:   s,
		rating: r,
	}
	return
}

// func main() {
// 	fmt.Printf("%+v", getMovie("xyz", 3.5))
// }

// ---------------------- Another question
func increaseRating(m *Movie) {
	(*m).rating += 1.0
}

// func main() {
// 	mov := getMovie("xyz", 2.0)
// 	increaseRating(&mov)
// 	fmt.Printf("%+v", mov)
// }

// --------------------- Another question -----need to understand

// func main() {
// 	mov := getMovie("xyz", 2.1)
// 	mov1 := getMovie("abc", 3.3)
// 	movies := make([]Movie, 5)
// 	movies = append(movies, mov)
// 	movies = append(movies, mov1)
// 	for _, value := range movies {
// 		fmt.Println(value)
// 	}
// }

// --------------------- Another question
func main() {
	mov := Movie{"xyz", 2.1}
	mov1 := Movie{"abc", 2.1}
	if mov.rating == mov1.rating || mov != mov1 {
		fmt.Println("condition met")
	} else if mov.rating == mov1.rating {
		fmt.Println("condition_2 met")
	}
}

// Output
// condition met
