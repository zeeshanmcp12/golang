package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Working with JSON. Encoding/Decoding")
	// EncodingJson()
	// DecodeJSON()
	// DecodeJSONFromFile()
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serverHome).Methods("GET")

}

func greeter() {
	fmt.Println("Hello from golang!")
}

func serverHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>Welcome to A Cloud Techie!</h1>"))

}

func EncodingJson() {

	userProfile := []User{
		{"Muhammad Zeeshan", 32, "+", "zeeshan@dev.io"},
		{"Hussain Raza", 28, "+", "hussain@dev.io"},
		{"Abdullah Shafiq", 25, "+", "raza@dev.io"},
	}

	finalJson, err := json.MarshalIndent(userProfile, "", "\t")
	CheckNilErr(err)

	fmt.Printf("%s\n", finalJson)

}

func DecodeJSON() {
	jsonData := []byte(`
		{
			"fullname": "Muhammad Zeeshan",
			"age": 32,
			"email": "zeeshan@dev.io"
		}
	`)

	isJsonValid := json.Valid(jsonData)

	var jsonStruct User

	if isJsonValid {
		json.Unmarshal(jsonData, &jsonStruct)
		fmt.Printf("%#v", jsonStruct)

	} else {
		fmt.Println("JSON IS NOT VALID")
	}

}

func DecodeJSONFromFile() {
	file, err := ioutil.ReadFile("./readjson.json")

	CheckNilErr(err)

	jsonData := []byte(file)

	var jsonStruct FileData

	validJson := json.Valid(jsonData)

	if validJson {
		fmt.Printf("JSON is Valid!")
		json.Unmarshal(jsonData, &jsonStruct)
		fmt.Printf("%#v\n", jsonStruct)
	} else {
		fmt.Println("JSON IS NOT VALID!")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData, &myOnlineData)

	for index, val := range myOnlineData {
		fmt.Printf("Key is %v Value is %v\n", index, val)
	}

	// Another case where you want to add data to key value
	// var myOnlineData map[string]interface{}
	// json.Unmarshal(jsonDataFromFile, &myOnlineData)
	// // fmt.Printf("%#v\n", myOnlineData)

	// for k, val := range myOnlineData { //k = key, val = value
	// 	fmt.Printf("Key: %v\tValue: %v\tType: %T\n", k, val, val)
	// }

}

func CheckNilErr(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	Name        string `json:"fullname"`
	Age         int    `json:"age"`
	Password    string `json:"-"`
	ContactInfo string `json:"email"`
}

type FileData struct {
	Name        string `json:"fullname"`
	Age         int    `json:"age"`
	ContactInfo string `json:"email"`
}
