package main

import "fmt"

func main() {
	fmt.Println("Starting maps in golang!")
	abbreviations := make(map[string]string)

	abbreviations["JS"] = "JavaScript"
	abbreviations["PS"] = "PowerShell"
	abbreviations["OS"] = "OperatingSystem"

	fmt.Println(abbreviations)
	fmt.Println("JS stands for:", abbreviations["JS"])

	delete(abbreviations, "PS")
	fmt.Println(abbreviations)

}
