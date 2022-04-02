package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Course Model
type Course struct {
	Courseid   string  `json:"courseid"`
	Coursename string  `json:"coursename"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

// Author Model
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware, helper - file
func (c Course) isEmpty() bool {
	return c.Courseid == "" && c.Coursename == ""

}

func main() {

}

// Controller - HomePage
// Controller is actually a route
func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to A Cloud Techie!</h1>"))

}

// Controller - Get All Course
func getAllData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	// Set Header first
	w.Header().Set("Content-Type", "application/json")

	// Create json encoder to encode the json data. In our case, it is course from fake db
	json.NewEncoder(w).Encode(courses)

}

// Controller - Get one course by ID
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course by id")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// Loop through the course,
	//find matching id,
	//return the response
	for _, course := range courses {
		// find matching id
		if course.Courseid == params["id"] {
			json.NewEncoder(w).Encode(course)
			fmt.Printf("Type of params is %T\n", params)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id!")
	return

}
