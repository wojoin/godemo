package main

import (
	"flag"
	"fmt"
	"time"
)

var conf = flag.String("c","log.conf","log configuration")
var logname = flag.String("n","logname","logname of configuration")
var pid = flag.Int64("p", 1,"process id")
var loglevel = flag.String("level","info","log level")
var period = flag.Duration("period",time.Second * 10, "the duration of output log")

func init() {
	flag.Parse()
}

func main() {
	fmt.Println("conf file:", *conf)
	fmt.Println("logname:", *logname)
	fmt.Println("pid:", *pid)
	fmt.Println("log level:", *loglevel)
	fmt.Println("period:", *period)
}