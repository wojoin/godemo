package main

import (
	"log"
	"strings"
)

func main(){
	log.Println("OK")
	s := "autp/private/Job/123"
	s = strings.ToLower(s)
	log.Println("low case :", s)

	if ok := strings.Contains(s, "job"); !ok{
		return
	}
	n := strings.LastIndex(s,"/")
	subresid := s[n+1:]
	log.Println("sub resource id :", subresid)
}
