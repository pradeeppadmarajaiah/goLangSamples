package main

import (
	"fmt"
)

//Struct has one extra byte memory for padding
type User struct {
	id      int
	name    string
	isAdmin bool
}

func main() {

	var user1 User
	fmt.Printf("User1 value is : %v\n", user1)

	user2 := User{}
	fmt.Printf("User2 value is : %v\n", user2)

	user3 := User{1, "Pradeep", false}
	fmt.Printf("User3 value is : %v\n", user3)

	user4 := User{id: 2, name: "Hitesh"}
	fmt.Printf("User4 value is : %v\n", user4)

	employee1 := struct {
		id      int
		company string
	}{id: 1, company: "IBM"}

	fmt.Printf("Employee1 company is : %v\n", employee1.company)

}
