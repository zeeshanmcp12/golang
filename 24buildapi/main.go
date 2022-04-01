package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
var course []Course

// middleware, helper - file
func (c Course) isEmpty() bool {
	return c.Courseid == "" && c.Coursename == ""

}

func main() {

}

// Controller
// HomePage

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to A Cloud Techie!</h1>"))

}

// Controller - Get All Course
func getAllData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(course)

}
