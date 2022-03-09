package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Learning Go Modules!")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET") //here we are using reference of function serveHome and not the actual function.
	fmt.Println("Listening on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func serveHome(w http.ResponseWriter, r *http.Request) {
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
