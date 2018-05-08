package main

import (
	"fmt"
)

func main() {
	//helloSendAndRecv()
	SignalAck()
}

func helloSendAndRecv() {
	ch := make(chan string)
	go func() { ch <- "hello" }()
	v := <-ch
	fmt.Println(v)
}

func SignalAck() {
	ch := make(chan string)

	go func() {
		fmt.Println(<-ch)
		ch <- "done"

	}()

	ch <- "started"
	fmt.Println(<-ch)

}
