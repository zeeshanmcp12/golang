package main

import "fmt"

func main() {
	fmt.Println("Writing methods in golang!")

	platformUser := User{"Zeeshan", 31, true}
	// fmt.Printf("User details are: \n%+v\n%d\n%t\n ", platformUser.Name, platformUser.Age, platformUser.IsActive)
	fmt.Printf("User details are:%+v\n", platformUser)

	platformUser.GetUserName()
	fmt.Println("Username has been updated!")
	platformUser.UpdateName()
}

type User struct {
	Name     string
	Age      int
	IsActive bool
}

func (u User) GetUserName() {
	fmt.Println("Username: ", u.Name)

}

func (u User) UpdateName() {
	// updatedName := u.Name{""}
	u.Name = "Abdullah"
	fmt.Println("Updated name: ", u.Name)

}
