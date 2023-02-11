package main

import "fmt"

type Video struct {
	name   string
	rating float64
}

func getVideo(n string, r float64) (v Video) {
	v = Video{
		name:   n,
		rating: r,
	}
	return
}

func increaseRating(v *Video) {
	(*v).rating += 1.0
}

func main() {
	fmt.Println("Struct lab questions!")
	vid := getVideo("Naat", 4.0)
	fmt.Printf("Type of vid -> %T\n", vid)
	fmt.Printf("Address of vid -> %v\n", &vid)
	fmt.Printf("Value of vid -> %+v\n", vid)

	var num1 int = 10
	fmt.Println(num1)
	numChanged := &num1
	*numChanged = 20
	fmt.Println(num1)

	increaseRating(&vid)

	fmt.Printf("Value -> %+v", vid)

}
