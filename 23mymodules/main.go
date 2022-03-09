package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Learning Go Modules!")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	fmt.Println("Listening on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Learning Modules in golang</h1>"))

}
