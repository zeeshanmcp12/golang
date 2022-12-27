package main

import "fmt"

func main() {
	fmt.Println("Build an API in golang!")
}

// Steps to create an API in golang!
// 1- Create a model (it will be struct in this case)
// 2- Create a fake db (a slice of model)
// 3- Helper function (or methods)

// Course Model - will go into file later on
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

// Author Model - will go into file later on
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// Helper method
func (c Course) helper() bool {
	return c.CourseId == "" && c.CourseName == ""
}
