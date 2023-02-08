package main

type Struct1 struct {
	number int
}

type Struct2 struct {
	number int
}

// It will give compile time error because struct itself is a type hence both structs are different.
// func main() {
// 	fmt.Println("Compare struct in golang!")

// 	structVar := Struct1{number: 5}
// 	structVar1 := Struct2{number: 5}

// 	if structVar == structVar1 {
// 		fmt.Println("Both structs are same!")
// 		// It will give compile time error because struct itself is a type hence both structs are different.
// 	}
// }

func main() {
	str1 := Struct1{number: 5}
	str2 := Struct1{number: 6}
	str3 := Struct1{number: 5}

}
