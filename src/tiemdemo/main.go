package main

import (
	"fmt"
	"time"
)

func main() {
		exp := time.Now().Unix()
		fmt.Println("now:",time.Now())
		fmt.Println("exp:",exp)
		fmt.Println("expiration",time.Unix(exp , 0).String())
		fmt.Println("Hello, 世界")

		timeLocal := time.FixedZone("CST", 3600*8)
		time.Local = timeLocal
		localTime := time.Now().Local()
		fmt.Println("localTime",localTime)
		fmt.Println("unix timestamp",time.Now().Unix())

}
