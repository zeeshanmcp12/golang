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
	fmt.Println("Struct and pass it by reference")
	vid := getVideo("nothing", 3.4)

	fmt.Printf("%+v\n", vid)
	increaseRating(&vid)
	fmt.Printf("Rating increased: %+v\n", vid)

	videos := make([]Video, 5)
	// fmt.Printf("%+v", videos)

	vid2 := getVideo("nothing", 4.0)

	videos = append(videos, vid)
	videos = append(videos, vid2)

	for _, val := range videos {
		fmt.Println(val)
	}

	updateVid := &videos
	*updateVid = append(*updateVid, Video{"Test1", 2.2})
	fmt.Println("New video added")
	for _, val := range videos {
		fmt.Println(val)
	}
}
