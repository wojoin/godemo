package main

import (
	"context"
	"fmt"
	"log"

	"net"
	"net/http"
	"os"
	"time"
)

var logg *log.Logger

func handler() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			time.Sleep(1 * time.Second)
			select {
			case <-ctx.Done():
				logg.Printf("done")
				return
			default:
				logg.Printf("work")
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	logg.Printf("cancel")
	cancel()
}

type Handler1 func(int)

func handler1(id int) {
	fmt.Printf("handle %d\n", id)
}

func main() {
	logg = log.New(os.Stdout, "INFO: ", log.Ltime)
	handler()

	Handler1(handler1)(42)

	logg.Printf("down")
}
