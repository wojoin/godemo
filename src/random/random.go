package main

import (
	"fmt"
	"math/rand"
	"time"
)

const LENGTH int = 10000000

func main() {

	//
	// 全局函数
	//

	rand.Seed(time.Now().Unix())

	var randNumber []int
	var index int

	for index = 0; index < LENGTH; index++ {
		randNumber = append(randNumber, rand.Intn(20000))
	}

	//fmt.Println(rand.Intn(200))

	//	for _, val := range randNumber {
	//		fmt.Println("val: ", val)
	//	}

}
