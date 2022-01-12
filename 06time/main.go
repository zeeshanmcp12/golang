package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	presentTime := time.Now()
	fmt.Println(presentTime.Format("02-01-2006 Monday"))

}
