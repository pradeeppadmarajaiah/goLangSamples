package main

import (
	"fmt"
)

type user struct {
	name string
	age  int
}

func (u *user) print() {
	fmt.Printf("name is %v and age is %v", u.name, u.age)
}

type employee struct {
	user     //Embedding with out calling a reference
	location string
}

type member struct {
	userDetail user // With out Embedding. Just calling a reference
	location   string
}

func main() {

	user1 := user{"Pradeep", 29}
	employee := employee{user: user1, location: "India"}

	employee.user.print()

	// or
	employee.print()
}
