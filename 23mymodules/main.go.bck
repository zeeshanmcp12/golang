package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Learning about mod in golang")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/google", serveGoogle).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}

func greeter() {

	fmt.Printf("Sample app listening at http://localhost:8000")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Thank you for visiting acloudtechie.com</h1>"))
	w.WriteHeader(http.StatusOK)

}

func serveGoogle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(readFile()))

}

func readFile() []byte {
	data, err := ioutil.ReadFile("../funExamples/site_google.html")
	CheckNilErr(err)
	return data
}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}

}
