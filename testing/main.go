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
	listenAddr := ":4000"
	fmt.Printf("About to Listen on %s. Go to http://127.0.0.1%s\n", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, r))

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
