package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func main() {
	secwebsocketkey := "dGhlIHNhbXBsZSBub25jZQ=="
	guid := "258EAFA5-E914-47DA-95CA-C5AB0DC85B11"
	key := secwebsocketkey + guid

	data := []byte(key)
	fmt.Println("key:", key)
	fmt.Printf("hash(key) = %x\r\n", sha1.Sum(data))
	sum := sha1.Sum(data)

	secwebsocketaccept := base64.StdEncoding.EncodeToString(sum[:])
	fmt.Println("Sec-WebSocket-Accept:", secwebsocketaccept)

}
