package main4

import (
	"errors"
	"fmt"
)

func getPaintValue(width, height float64) (value float64, err error) {
	if width < 0 {
		return 0, fmt.Errorf("width value %.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("width value %.2f is invalid", width)
	}
	area := width * height
	value = area / 10.0
	return value, nil

}

func main4() {

	err := errors.New("this is an error")

	//fmt can detect err.Error() or err as same
	fmt.Println(err.Error())
	fmt.Println(err)

	getPaintValue, err := getPaintValue(-23.23, 34.45)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Required paint in liters is %.2f \n", getPaintValue)
	}

}
