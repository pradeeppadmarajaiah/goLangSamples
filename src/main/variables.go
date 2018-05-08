package main

import (
	"fmt"
)

var a int = 34
var b string
var c bool
var d float64

//not allowed
//g:=23

func main() {

	g := 23
	f := float64(a)
	fmt.Println(g, "  ", f)
}
