package main5

import (
	"fmt"
)

func triple(input int) int {
	return input * input * input
}

func tripleP(input *int) {
	*input = *input * *input * *input
}

func main5() {
	age := 30
	ageAddress := &age

	fmt.Printf("variable memory address is %v . Variable value is %v \n", ageAddress, *ageAddress)

	value := 2

	triple := triple(value)

	fmt.Printf("Without Pointers : value is %v . Tripled value is %v\n", value, triple)

	valueAdress := &value

	tripleP(valueAdress)
	fmt.Printf("With Pointers : value is %v ", value)

}
