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
	fmt.Println("Struct Lab Questions!")
	vid := getVideo("nothing", 3.4)
	fmt.Printf("%+v", vid)
	increaseRating(&vid)
	fmt.Printf("\n%+v", vid)

}
