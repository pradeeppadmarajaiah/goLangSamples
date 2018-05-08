package main6

import (
	"encoding/json"
	"fmt"
)

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

type Order struct {
	Id         int
	Name       string
	TotalPrice float64
}

func main6() {

	intV, err := json.Marshal(23)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(intV))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	var res response2
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Page)
	fmt.Println(res.Fruits[0])

	order1 := Order{1, "Order1", 12.345}
	b, _ := json.Marshal(order1)

	fmt.Printf("Marshalled value is %v", string(b))

	var m Order
	err1 := json.Unmarshal(b, &m)

	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println(m)

}
