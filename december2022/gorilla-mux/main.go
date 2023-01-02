package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Working with mux router")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/profile", profile).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Server is listening on http://localhost:4000")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to https://acloudtechie.com</h1>"))

}

func profile(w http.ResponseWriter, r *http.Request) {
	jsonData := []Profile{
		{"Zeeshan", 32, "acloudtechie.com"},
		{"Abdullah", 30, "abc.com"},
	}

	finalJSON, _ := json.MarshalIndent(jsonData, "", "\t")

	fmt.Println(string(finalJSON))

	w.Write([]byte(finalJSON))

}

type Profile struct {
	Name    string
	Age     int
	Website string
}
