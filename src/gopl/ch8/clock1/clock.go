package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main(){
	listener, err := net.Listen("tcp4", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for{
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		//deal with only one client at a time,
		// but adding go statement causes each call to run its own goroutine
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer  conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}

		time.Sleep(time.Second * 1)
	}
}
