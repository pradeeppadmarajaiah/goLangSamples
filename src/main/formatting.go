package main3

import (
	"fmt"
)

func getPaintValue(width, height float64) {
	area := width * height
	//fmt.Printf("Total paint neeeded for the height : %f  & width : %f is %f \n", height, width, area/10.0)

	fmt.Printf("Total paint neeeded is : %0.3f .which has height : %f and width : %f \n", area/10.0, height, width)
}

func main3() {
	fmt.Printf("My name is %s \n", "Pradeep")
	fmt.Printf("My name is %33s \n", "Pradeep")
	fmt.Printf("My Salary is %.2f \n", 1.199999999)

	width, height := 23.34, 23.4545

	fmt.Printf("area is %.2f \n", height*width)

	getPaintValue(12.23, 23.23)

}
