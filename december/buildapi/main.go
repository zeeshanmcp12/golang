package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Build an API in golang!")
}

// Steps to create an API in golang!
// 1- Create a model (it will be struct in this case)
// 2- Create a fake db (a slice of model)
// 3- Helper function (or methods)

// Course Model - will go into file later on
// Author is a type that we created below and since it is not created at the moment so we are passing reference of it
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
// This fake db will be a slice of type "Course" that we created above
var courses []Course

// Helper method or sometimes these are called middleware
// These methods usually go inside a separate file
// We are passing pointer of course here that's why wrote like this *Courses
func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

// Create controllers - this will also go inside separate file
// Controller 1 - serveHome
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to acloudtechie.com</h1>"))
}

// Controller 1 - getAllCourses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")

	// Here, we are setting up header using Header method and then Set method
	w.Header().Set("Content-Type", "application/json")

	// Here, we encoding data into json. So NewEncoder requires response write which is w in our case
	// .Encode is a method which will encode data into json format whatever we will send in our response
	// courses -> is a slice of type Course
	json.NewEncoder(w).Encode(courses)

}
