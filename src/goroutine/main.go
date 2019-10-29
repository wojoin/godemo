package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// goroutine exit
// 4 ways

//func main(){
//
//	var inCh = make(chan int)
//
//	go func(in <-chan int) {
//		// Using for-range to exit goroutine
//		// range has the ability to detect the close/end of a channel
//		for x := range in {
//			fmt.Printf("Process %d\n", x)
//		}
//	}(inCh)
//
//
//	for i:=0;i<10;i++{
//		inCh <- i
//	}
//	close(inCh)
//
//	var in1 = make(chan int)
//	var in2 = make(chan int)
//
//	go func(in1 chan int, in2 chan int) {
//		// in for-select using ok to exit goroutine
//		for {
//			select {
//			case x, ok := <-in1:
//				if !ok {
//					in1 = nil
//				}
//				fmt.Printf("recv in1 %d\n", x)
//				// Process
//			case y, ok := <-in2:
//				if !ok {
//					in2 = nil
//				}
//				fmt.Printf("recv in2 %d\n", y)
//				// Process
//			//case <-t.C:
//			//	fmt.Printf("Working, processedCnt = %d\n", processedCnt)
//			}
//
//			// If both in channel are closed, goroutine exit
//			if in1 == nil && in2 == nil {
//				return
//			}
//		}
//	}(in1, in2)
//
//	for i:=0;i<10;i++{
//		in1 <- i
//		in2 <- 2*i
//	}
//
//	close(in1)
//	close(in2)
//
//}

type Error struct {
	message string
}

func (e Error) Error() string {
	return e.message
}

func goroutineFunc() {
	var wg sync.WaitGroup
	waitGroupLength := 8
	errChannel := make(chan error, 1)

	// Setup waitgroup to match the number of go routines we'll launch off
	wg.Add(waitGroupLength)
	finished := make(chan bool, 1) // this along with wg.Wait() are why the error handling works and doesn't deadlock

	for i := 0; i < waitGroupLength; i++ {

		go func(i int) {
			fmt.Printf("Go routine %d executed\n", i+1)

			// Sleep for the time needed for each other go routine to complete.
			// This helps show that the program exists with the last go routine to fail.
			// comment this line if you want to see it fail fast
			time.Sleep(time.Duration(waitGroupLength - i))

			time.Sleep(0) // only here so the time import is needed

			// comment out the following 3 lines to see what happens without an error
			// Note, the channel has a length of one so the last go routine to error
			// will always be the last error.

			if i%4 == 1 {
				errChannel <- Error{fmt.Sprintf("Errored on routine %d", i+1)}
			}

			// Mark the wait group as Done so it does not hang
			wg.Done()
		}(i)
	}

	// Put the wait group in a go routine.
	// By putting the wait group in the go routine we ensure either all pass
	// and we close the "finished" channel or we wait forever for the wait group
	// to finish.
	//
	// Waiting forever is okay because of the blocking select below.
	go func() {
		wg.Wait()
		close(finished)
	}()

	// This select will block until one of the two channels returns a value.
	// This means on the first failure in the go routines above the errChannel will release a
	// value first. Because there is a "return" statement in the err check this function will
	// exit when an error occurs.
	//
	// Due to the blocking on wg.Wait() the finished channel will not get a value unless all
	// the go routines before were successful because not all the wg.Done() calls would have
	// happened.
	select {
	case <-finished:
	case err := <-errChannel:
		if err != nil {
			fmt.Println("error ", err)
			return
		}
	}

	fmt.Println("Successfully executed all go routines")
}


func fetchAll() error {
	var N = 4
	quit := make(chan bool)
	errc := make(chan error) // err occur
	done := make(chan error) // nil error
	for i := 0; i < N; i++ {
		go func(i int) {
			// dummy fetch
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

			err := error(nil)

			if rand.Intn(4) == 0 {
				err = fmt.Errorf("goroutine %d's error returned", i)
			}
			ch := done // we'll send to done if nil error and to errc otherwise
			if err != nil {
				ch = errc
			}
			select {
			case ch <- err:
				return
			case <-quit:
				return
			}
		}(i)
	}



	count := 0
	for {
		select {
		case err := <-errc:
			close(quit)
			return err
		case <-done:
			count++
			if count == N {
				return nil // got all N signals, so there was no error
			}
		}
	}
}

func main() {
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(fetchAll())

	goroutineFunc()
}