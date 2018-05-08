package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int64
var mutex sync.Mutex

func main() {
	runtime.GOMAXPROCS(1)
	const grs = 2

	var wg sync.WaitGroup
	wg.Add(grs)

	for index := 0; index < grs; index++ {

		go func() {
			incCounter()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
}

func incCounter() {
	for index := 0; index < 2; index++ {
		mutex.Lock()
		{
			value := counter
			value++
			counter = value
		}
		mutex.Unlock()
	}

}
