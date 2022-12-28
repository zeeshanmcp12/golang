package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Build an API in golang!")

	r := mux.NewRouter()
	r.HandleFunc("/", homepage)
	r.HandleFunc("/course", getcourses)

	log.Fatal(http.ListenAndServe(":4000", r))
}

// Course Model
// Author Model
// Fake DB
// Helper methods
// 	1- that will check if course name and id is unique
// 	2- that will print server address in stream
// Controller - Router
// 	1- Serve Home / Homepage
// 	2- Get All Courses
//  3- Get Course by Id

// Course Model
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

// Author Model
type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// Helper method that will check if course name and id is unique
func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func greeter() {
	fmt.Println("Listening on http://localhost:4000")
}

// Controller - Router
// 	1- Serve Home / Homepage
func homepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Hussain's Home</h1>"))
}

func getcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")

	// Here, we are setting up header using Header method and then Set method
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
