package main

import "fmt"

// First we need to create models
// These models should be in separate files but for the sake of simplicity and understanding we are creating in same file
// Course Model
type Course struct {
	CourseId   string
	CourseName string
	Price      int
	Website    string
	Author     *Author // We are defining our custom struct type here, so it was created as Pointer because we cannot pass as reference.
}

// Author model
// This model is referenced in Course struct
type Author struct {
	FirstName string
	LastName  string
}

// Now, Let's create helper method which will check if course id and course name is empty or not
// Since we've mentioned Course as pointer hence we will have access on all the properties in Course struct
func isCourseEmpty(c *Course) bool {
	return c.CourseId != "" && c.CourseName != ""
}

func main() {
	fmt.Println("Building API in golang!")
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
