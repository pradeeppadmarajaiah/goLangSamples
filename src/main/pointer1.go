package main

import (
	"fmt"
)

func main() {
	count := 10

	fmt.Printf("Before : Count value is : %v.\n  Count address is : %v\n", count, &count)

	incrementCounter(count)

	fmt.Printf("After : Count value is : %v.\n  Count address is : %v\n", count, &count)

	increment(&count)
	fmt.Printf("After pointing via address: Count value is : %v.\n  Count address is : %v\n", count, &count)

	employee1 := getEmployee()

	fmt.Printf("Default employee details %v . value is %v", employee1, *employee1)

	*employee1 = Employee{"Deep", "USA"}
	fmt.Printf("After: employee details %v . value is %v\n", employee1, *employee1)

}

func incrementCounter(count int) {
	count = count + 5
}

func increment(count *int) {
	fmt.Printf("value of count is %v. value of *count is %v. value of &count ", count, *count, &count)
	*count = *count + 5
}

//Not recommended one
func getEmployee() *Employee {
	user := &Employee{"Pradeep", "India"}
	return user
}

type Employee struct {
	name     string
	location string
}
