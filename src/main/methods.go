package main1

import (
	"fmt"
)

//MyInt : input
type MyInt int

//Weight : hello
type Weight struct {
	width, height int
}

func (a MyInt) sum() MyInt {
	return a + MyInt(10)
}

func (a *MyInt) square() {
	*a = (*a) * (*a)
}

func main1() {
	var input MyInt = 2
	fmt.Println(input)
	input.square()
	fmt.Println(input)

	fmt.Println(input.sum())

	fmt.Println("Welcome to deepus world")
}
