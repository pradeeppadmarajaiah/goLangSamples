package main1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//https://tour.golang.org/methods/9

type Math interface {
	sum() int
}

//Myint : Hello
type Myint int

func (value Myint) sum() int {
	return int(value) + int(value)
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println(v, " : is a Value of type : int")
	case string:
		fmt.Println(i, " : is a Value of type : ", v)
	default:
		fmt.Println("Value is a different type")

	}
}

func main1() {

	fileInfo, error := os.Stat("interface.go")
	if error != nil {
		log.Fatal(error)

	}
	fmt.Println(fileInfo.Size())

	var a Math
	myInt := Myint(20)

	a = myInt

	fmt.Println(a.sum())

	var z Math = Myint(30)
	fmt.Println(z.sum())

	checkType(23)

	fmt.Print("Enter the name : ")
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')

	//input := "Pradeep"
	//error := "hello"

	if error != nil {
		log.Fatal(error)
	}

	input = strings.TrimSpace(input)

	grade, error := strconv.ParseFloat(input, 64)

	var status string
	if error != nil {
		log.Fatal("Entered input is not correct")
	}
	if grade > 35 {
		status = ("Passed")
	} else {
		status = ("failed")
	}
	fmt.Println(status)

}
