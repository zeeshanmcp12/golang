package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Build an API in golang!")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/getallcourses", getAllCourses).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

// Steps to create an API in golang!
// 1- Create a model (it will be struct in this case)
// 2- Create a fake db (a slice of model)
// 3- Helper function (or methods)

// Further dig down of action items
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

func greeter() {
	fmt.Println("Listening on http://localhost:4000")
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
	// w -> is response writer - whatever the response we will be sending to request.
	// courses -> is a slice of type Course
	json.NewEncoder(w).Encode(courses)

}

// Controller 2 - getOneCourse

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")

	// Here, we are setting up header using Header method and then Set method
	w.Header().Set("Content-Type", "application/json")

	// This time, user will request for course and send course id in the request
	// So, we need to grab that course id from request that user will send (or do request in url)
	// To grab the request id we need to use mux library which will have access on query params
	// grab id from request
	params := mux.Vars(r)

	fmt.Printf("Type of params is %T\n", params)

	// Once we found out the id, we need to perform following action:
	// loop through the array (Course fake db)
	// Find the matching id
	// Return the response

	// Loop through the courses array (fake db)
	for _, course := range courses {

		// Find the matching id
		if course.CourseId == params["id"] {

			// Return the response
			// We can return course name if course id is found or we can return all details
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id %v", params["id"])
	return

}

// Controller 3 - createOneCourse
// Action plan
// 1- Set Header
// 2- Check if body is empty
// 3- Create course variable of type Course
// 4- Decode JSON to check if there is some data or only blank curly braces
// 5- Check if json is empty and having only {} "curly braces"
// 6- Generate unique id (course id)
// 7- Convert int to string using string.itoa
// 8- Seed the data in database which is Course fake db in our case
// 		8.1- Append the course in Course fake db

func createOneCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Body is Empty")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data in JSON body!")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}
