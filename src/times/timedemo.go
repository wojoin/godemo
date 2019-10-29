package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var str = "2019-05-29"
	arr := strings.Split(str,"-")
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	year, _ := strconv.Atoi(arr[0])
	month, _ := strconv.Atoi(arr[1])
	day,_ := strconv.Atoi(arr[2])
	//t := time.Date(strconv.Atoi(arr[0]), time.Month(arr[1]), strconv.Atoi(arr[2]))
	//t, _ := time.Parse(shortForm, "2019-05-29")
	//fmt.Println(t)

	t := time.Date(year, time.Month(month), day, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
	fmt.Printf("Go launched at %s\n", t.Unix())

	fmt.Println(time.Now().Format("2006-01-02"))


}
