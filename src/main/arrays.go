package main5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main5() {
	fmt.Println("Welcome to arrays world!!!!!!")

	numbers := [5]int{12, 12, 43, 34, 34}
	var numberList [5]int
	numberList[0] = 1
	numbers = numberList

	fmt.Println("Numbers : ", numbers)

	colors := []string{"red", "blssue", "green", "black", "orange", "white", "grey", "marron", "pink", "yellow", "green"}

	fmt.Printf("color length is %v", len(colors))

	for index := 0; index < len(colors); index++ {
		fmt.Printf("%v  | %s \n", index, colors[index])
	}

	fmt.Println("Using range in for loop")
	for index, color := range colors {
		fmt.Printf("%v  | %s \n", index, color)
	}

	sum := 0
	for _, number := range numbers {

		sum += number
	}
	fmt.Printf("Avg of num is %v", sum)

	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal("Error while reading the file!")
	}

	var valueList [3]float64
	i := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		valueList[i], _ = strconv.ParseFloat(scanner.Text(), 64)
		i += 1
	}

	fmt.Printf("\nData from file %v", valueList)

	err = file.Close()
	if err != nil {
		log.Fatal("Error while closing the file")
	}
	if scanner.Err() != nil {
		log.Fatal("Error while scanning the file")

	}
}
