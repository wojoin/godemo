package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan int = make(chan int, 10)
	go send(c)
	go recv(c)
	time.Sleep(6 * time.Second)
	close(c)
}

func send(c chan<- int) {
	for i := 0; i < 5; i++ {
		fmt.Println("send ready...")
		time.Sleep(1 * time.Second)
		c <- i
		fmt.Println("send ok")

	}
}

func recv(c <-chan int) {
	for v := range c {
		fmt.Printf("recv %d\n", v)
	}
}
