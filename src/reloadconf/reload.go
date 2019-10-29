package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	Mode string `json:"Mode"`
}

var (
	config     *Config
)

func loadConfig() {
	f, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println("[TEST_SIGUSR] Load config: ", err)
	}
	err = json.Unmarshal(f, &config)
	if err != nil {
		fmt.Println("[TEST_SIGUSR] Para config failed: ", err)
	}
}

func init() {
	loadConfig()
	fmt.Println("[TEST_SIGUSR] Load config;Mode: ", config.Mode)
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGUSR2)
	go func() {
		for {
			<-s
			loadConfig()
			fmt.Println("[TEST_SIGUSR] ReLoad config;Mode: ", config.Mode)
		}
	}()

}
func main() {
	q := make(chan os.Signal, 1)
	signal.Notify(q, syscall.SIGINT)
	for {
		select {
		case <-q:
			fmt.Println("[TEST_SIGUSR] ---ibye----")
			os.Exit(1)
		default:
		}
		fmt.Println("[TEST_SIGUSR] ---waiting to reload----")
		time.Sleep(time.Second * 2)

	}

}