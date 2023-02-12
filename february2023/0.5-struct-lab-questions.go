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

func modifyName(v *Video) {
	(*v).name = "Something new"
}

func main() {
	fmt.Println("Struct Lab Questions!")
	vid := getVideo("nothing", 3.4)
	fmt.Printf("%+v\n", vid)

	// Update only rating using function (or Pass struct by reference)
	increaseRating(&vid)
	fmt.Printf("%+v\n", vid)

	mov := &Video{"abc", 5.5}
	fmt.Println(mov)

	// Update only name using dereferencing
	(*mov).name = "bcd"
	fmt.Println(mov)

	// Update only name using function (or Pass struct by reference)
	modifyName(&vid)
	fmt.Printf("%+v\n", vid)

}
