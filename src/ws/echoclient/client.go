package main

import (
	"flag"
	"time"

	//"fmt"
	"log"
	"net/url"

	//"net/url"
	"os"
	"os/signal"

	//"time"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func init() {
	log.SetPrefix("TRACE: ")
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags)
}

func main() {
	flag.Parse()
	log.Printf("address: %s\n", *addr)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s\n", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial: ", err)
	}

	defer conn.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			msgtype, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: message type %d, message %s", msgtype, message)

		}
	}()

	// start timer
	ticket := time.NewTicker(time.Second)
	defer ticket.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticket.C: // receive ticket from time channel C
			err := conn.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Printf("write:", err)
				return
			}
			time.AfterFunc(5*time.Second, func() {
				done <- struct{}{}
				log.Println("done")
			})

		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
