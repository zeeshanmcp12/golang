package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Testing file")

	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/data", showData).Methods("GET")
	fmt.Println("Listening on port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))

}

func greeter() {
	fmt.Println("Hello from golang!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to A Cloud Techie!</h1>"))

}

func showData(w http.ResponseWriter, r *http.Request) {

	jsonData := []User{
		{"Zeeshan", 32, "acloudtechie@dev.co"},
		{"Abdullah", 30, "zeeshan@dev.co"},
	}

	fmt.Println(jsonData)

	finalJson, _ := json.MarshalIndent(jsonData, "", "\t")

	w.Write([]byte(finalJson))

}

type User struct {
	Name  string
	Age   int
	Email string
}
