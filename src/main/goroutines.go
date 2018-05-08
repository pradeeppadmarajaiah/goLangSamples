package main7

import (
	"fmt"
	"sync"
)

func printName(name string) {
	fmt.Printf("Hello %v", name)
}

func main7() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("Welcome of anonymous function")
		wg.Done()
	}()

	wg.Add(1)
	sayName := func() {
		fmt.Println("hello")
		wg.Done()
	}

	go sayName()

	wg.Wait()
	printName("Pradeep")
}
