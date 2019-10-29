package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := `(^1\d{10}$)`
	reg := regexp.MustCompile(pattern)

	fmt.Println(reg.MatchString("58616931990"))

}
