package main

import "fmt"

func main() {
	fmt.Println("Maps in golang!")

	// var codes map[string]string // map initialization with value throws runtime error because it's nil
	// codes["en"] = "English"     // panic: assignment to entry in nil map

	// Initializing and declaring map with values
	data := map[string]string{"Name": "Zeeshan", "Website": "acloudtechie.com"}

	fmt.Println(data)
	// fmt.Println("Key value pair in data:", len(data)) // 2 which is the length of elements

	// Add key value in map
	data["Location"] = "Pakistan"
	data["Dummy"] = "Value"

	// Update key value in map
	data["Website"] = "https://acloudtechie.com"

	// Delete key value pair
	delete(data, "Dummy")

	// Loop through map with range keyword
	for key, value := range data {
		fmt.Println(key, "=>", value)
	}

	// declaring and initializing map with make function
	// age := make(map[string]int)

	// Accessing a map
	fmt.Println("Website:", data["Website"])

	dummyMap := map[string]int{"en": 1, "ar": 2, "pk": 3}

	// getting a key
	// When we say getting a key in maps, it means getting a value which is associated with that key
	// to get a key in map:
	fmt.Println()
	value, found := dummyMap["en"]
	fmt.Println(found, value) //true 1

	value2, found2 := data["hh"]
	fmt.Println(found2, value2) // false <empty space>

	// truncate a map
	// using for loop with range keyword
	// for key, _ := range data {
	// 	delete(data, key)
	// }
	// fmt.Println("map truncated!", data)

	// Reinitialize empty map
	data = make(map[string]string)
	fmt.Println("Map truncated", data)

}
