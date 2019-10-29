package main

import (
	"fmt"
	//"io"
	"log"
	//"net/http"
	"time"

	"github.com/valyala/fasthttp"
)

func main() {
	// echo time, the web server
	var count int64
	//	timeHandler := func(w http.ResponseWriter, req *http.Request) {
	//		time.Sleep(time.Millisecond * 150)
	//		count++
	//		fmt.Println("host: ", req.Host)
	//		fmt.Println("path: ", req.URL.Path)
	//		fmt.Println("count: ", count)
	//		io.WriteString(w, time.Now().String())
	//	}

	listenAddr := "10.64.39.52:8080"

	// This function will be called by the server for each incoming request.
	//
	// RequestCtx provides a lot of functionality related to http request
	// processing. See RequestCtx docs for details.
	timeHandler2 := func(ctx *fasthttp.RequestCtx) {
		time.Sleep(time.Millisecond * 150)
		count++
		fmt.Println("count: ", count)
		fmt.Fprintf(ctx, "path: %q\n", ctx.Path())
		fmt.Fprintf(ctx, "time: %s\n", time.Now().String())
	}

	// Start the server with default settings.
	// Create Server instance for adjusting server settings.
	//
	// ListenAndServe returns only on error, so usually it blocks forever.
	if err := fasthttp.ListenAndServe(listenAddr, timeHandler2); err != nil {
		log.Fatalf("error in ListenAndServe: %s", err)
	}
	//fmt.Println("count: ", count)

	//go http.HandleFunc("/time", timeHandler2)
	//log.Fatal(http.ListenAndServe(":8080", nil))
}
