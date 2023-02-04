package main

import "fmt"

type Student struct {
	Name   string
	rollNo int
	marks  []int
}

func main() {

	var studentData Student
	fmt.Printf("%+v", studentData)

}
