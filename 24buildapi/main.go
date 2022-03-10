package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// First we need to create models
// These models should be in separate files but for the sake of simplicity and understanding we are creating in same file
// Course Model
type Course struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

// We are defining our custom struct type here, so it was created as Pointer because we cannot pass as reference.

// Author model
// This model is referenced in Course struct
type Author struct {
	FullName string `json:"name"`
	Website  string `json:"website"`
}

// fake DB
// We also need to create a fake db (only at this stage)
// this courses fake db is gonna be a slice of type Course (we defined above)
var courses []Course

// Now, Let's create helper method which will check if course id and course name is empty or not
// Since we've mentioned Course as pointer hence we will have access on all the properties in Course struct
// It's actually a method and not function
// It will return boolean value (true and false)
func (c *Course) isEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

// Controllers - file (mean goes into separate file)

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to acloudtechie.com</h1>"))

}

// This is to be understand again
func getAllVal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all cources")
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func main() {
	fmt.Println("Building API in golang!")
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
