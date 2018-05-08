package main

import (
	"fmt"
	"runtime"
	"sync"
)

func lowercase() {
	for index := 0; index < 1000; index++ {
		fmt.Println("LowerCase ", index)
	}

}

func uppercase() {
	for index := 0; index < 1000; index++ {
		fmt.Println("UpperCase ", index)
	}
}

func main() {

	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		lowercase()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		uppercase()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Completed all the go routines")

}
