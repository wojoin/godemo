package main

import "fmt"

func main() {
	natural := make(chan int)
	squares := make(chan int)

	// counter
	go func () {
		for i := 0;i < 100 ; i++ {
			natural <- i
		}
		close(natural)
	}()

	// squares
	go func() {
		for {
			value, ok := <- natural
			if !ok {
				break
			}
			squares <- value * value
		}
		close(squares)
	}()

	for {
		s, ok := <-squares
		if !ok{
			break
		}
		fmt.Println(s)
	}

	//close(natural)
	//close(squares)
}
