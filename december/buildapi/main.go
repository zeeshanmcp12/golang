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

	// Seeding database
	courses = append(courses, Course{CourseId: "2", CourseName: "JavaScript", CoursePrice: 1000, Author: &Author{FullName: "Abdullah", Website: "abc.com"}})
	courses = append(courses, Course{CourseId: "3", CourseName: "Golang Series", CoursePrice: 1500, Author: &Author{FullName: "Ahmed", Website: "xyz.com"}})

	// Setup routers
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course/{id}", createOneCourse).Methods("POST")

	// Listen to port 4000
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
	// return c.CourseId == "" && c.CourseName == ""

	// Previously we were returning true and false if course id and course name is empty
	// But now we are only returning course name
	// Because we want to generate the course id ourselve and not let user generate this
	return c.CourseName == ""
}

func greeter() {
	fmt.Println("Listening on http://localhost:4000")
}

// Create controllers - this will also go inside separate file
// Controller 1 - serveHome
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to acloudtechie.com</h1>"))
}

// Controller 2 - getAllCourses
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

// Controller 3 - getOneCourse

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

// Controller 4 - createOneCourse
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

	// Check if request body is empty
	// r.Body -> r is request which has access on Body.
	if r.Body == nil {
		// If body is empty then send response back
		json.NewEncoder(w).Encode("Body is Empty")
	}

	// Create course variable of type Course - The type (or struct) that we created before
	// This variable is to create an object so we can pass reference of it below
	var course Course

	// Decode JSON to see what's coming in (request of) JSON body.
	_ = json.NewDecoder(r.Body).Decode(&course)

	// Check if JSON body has no data
	if course.isEmpty() {
		// If body is empty then send response back
		json.NewEncoder(w).Encode("No data in JSON body!")
		return
	}

	// Generate unique id
	// Generate random numbers
	rand.Seed(time.Now().UnixNano())

	// Convert generated numbers from int to string using strconv package
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// append data (course) into courses -> this 'courses' is a slice we created at line number 60
	// this will be new course that will add into courses (slice)
	courses = append(courses, course)

	// Send response that new course has been added into courses
	fmt.Println("New Course has been added")
	json.NewEncoder(w).Encode(course)
	return

}

// Controller 5 - updateOneCourse
// Action plan
// 1- Set Header
// 2- Grab course id(s) because to find out the course that has to be updated we need an id of that course
// 3- Loop through the courses
// 4- Match the requested course id with all courses
// 5- Remove the course id when found
// 6- Update the course with my id (which is requested in request)
// 7- Send the response back (OR Craft the response using json.NewEncoder())

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// Grab the course id
	params := mux.Vars(r)

	// Loop through all courses
	for index, course := range courses {
		// Find the requested id (Match the course id)
		if course.CourseId == params["id"] {
			// Remove the course when course is found against provided id
			courses = append(courses[:index], courses[index+1:]...)

			// Create new object
			var course Course
			// Decode JSON to see what is coming inside JSON body
			_ = json.NewDecoder(r.Body).Decode(&course)

			// Update the course with new id
			course.CourseId = params["id"]

			// When course id updated then append it to all courses
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: send a response when id is not found
}

// Controller 6 - deleteOneCourse
// Action plan
// 1- Set Header
// 2- Grab course id(s) because to find out the course that has to be updated we need an id of that course
// 3- Loop through the courses
// 4- Match the requested course id with all courses
// 5- Remove the course id when found
// 6- Craft the response using json.NewEncoder()
func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course has been REMOVED!")
			break
		}
	}
}
