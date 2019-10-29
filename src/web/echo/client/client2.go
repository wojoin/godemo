package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	LENGTH int64  = 1000
	PORT   int    = 8080
	PATH   string = "/time"
)

var envoyhost = flag.String("envoyhost", "", "host deployed envoy")
var envoyport = flag.String("envoyport", "", "port listened by envoy")
var logger = flag.String("log", "", "port listened by envoy")

func main() {
	flag.Parse()

	c := &fasthttp.Client{}

	url := "http://10.64.39.52:8080/time"
	if *envoyhost != "" && *envoyport != "" {
		//url := "http://10.64.39.52:8080/time"
		url = "http://" + *envoyhost + ":" + string(*envoyport) + PATH
		fmt.Println("url: ", url)
	}

	fmt.Println("url: ", url)

	//	client := &http.Client{}

	//	reqest, err := http.NewRequest("GET", url, nil)
	//	if err != nil {
	//		panic(err)
	//	}

	//t := make(map[int]int64)

	var total int64

	for index := 0; index < int(LENGTH); index++ {
		start := time.Now()
		//response, _ := client.Do(reqest)
		//defer response.Body.Close()
		statusCode, body, err := c.Get(nil, "http://10.64.39.52:8080/time")
		if err != nil {
			log.Fatalf("Error when loading google page through local proxy: %s", err)
		}
		if statusCode != fasthttp.StatusOK {
			log.Fatalf("Unexpected status code: %d. Expecting %d", statusCode, fasthttp.StatusOK)
		}

		elapsed := time.Since(start)

		//		if !strings.HasPrefix(response.Status, "200") {
		//			fmt.Println("error, status code was not 200, ", response.Status)
		//		}
		//		if response.StatusCode != http.StatusOK {
		//			fmt.Println("error, status code was not 200, ", response.Status)
		//		}

		if *logger != "" {
			//			stdout := os.Stdout
			//			_, err = io.Copy(stdout, body)

			fmt.Println("result: ", string(body))
			fmt.Println("\telapsed: ", elapsed)
			fmt.Println("\tstatus: ", statusCode)
		}

		total += int64(elapsed)
	}
	fmt.Println("total: ", total)
	avg := total / LENGTH / 1000000
	fmt.Println("average: ", avg)
}
