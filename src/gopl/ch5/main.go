package main

import (
	"demo/src/gopl/ch5/links"
	"fmt"
	"log"
	"os"
)

func crawl(url string) []string {
	fmt.Println("url: ", url)
	list, err := links.Extract(url)
	if err != nil {
		log.Println(err)
	}
	for index, link := range list {
		log.Printf("index: %d, link: %s\n", index, link)
	}

	return list
}

func breadthFirst(extract func(url string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0  {
		items := worklist
		worklist = nil

		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, extract(item)...)
			}
		}
	}
}

func main() {
	//url := "http://www.baidu.com"
	breadthFirst(crawl, os.Args[1:])
}
