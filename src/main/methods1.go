package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("Sending email to %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {

	user2 := &user{"chandu", "chand@gmail.com"}
	user2.notify()

	user1 := user{"Pradeep", "pradeep@gmail.com"}
	//same as below
	//user1 := &user{"Pradeep", "pradeep@gmail.com"}
	user1.notify()
	//(*user1).notify()
	user1.changeEmail("deepu@gmail.com")
	user1.notify()

	users := []user{
		user{"hitesh", "hitesh@gmail.com"},
		user{"raj", "raj@gmail.com"},
	}

	for _, u := range users {
		u.notify()
	}
}
