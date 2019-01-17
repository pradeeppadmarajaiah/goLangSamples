package main

import (
	"fmt"
)

func main() {
	//helloSendAndRecv()
	// SignalAck()
	//SignalClose()
	closeRange()
}

//other way to close
func SignalClose() {
	ch := make(chan struct{})

	go func() {
		fmt.Println("signalClose() inprogress started")
		close(ch)
	}()

	fmt.Println("signalClose() inprogress started1")
	<-ch
	fmt.Println("signalClose() ended")
}

func helloSendAndRecv() {
	ch := make(chan string)
	go func() {
		//Receive ch <-
		ch <- "hello"

	}()
	//SEND <- ch
	v := <-ch
	fmt.Println(v)
}
func SignalAck() {
	ch := make(chan string)
	go func() {
		fmt.Println("bye " + <-ch)
		ch <- "done"
	}()
	ch <- "started"
	fmt.Println("hello " + <-ch)
}

func closeRange() {
	ch := make(chan int, 5)
	for index := 0; index < 5; index++ {
		ch <- index
	}
	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	for value := range ch {
		fmt.Println(value)
	}
}
