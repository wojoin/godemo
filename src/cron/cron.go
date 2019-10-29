package main

import (
	"fmt"
	"log"

	"time"

	"github.com/robfig/cron"
)

func task() {
	log.Println("execute per second")
}

const OneSecond = 1*time.Second + 10*time.Millisecond

func TestFuncPanicRecovery() {
	cronTask := cron.New()

	cronTask.Start()
	defer cronTask.Stop()
	cronTask.AddFunc("* * * * * ?", func() { log.Println("exec task one second") })

	for {
		select {
		case <-time.After(OneSecond):
			return
		}
	}
}

//func task2() {
//	timer1 := time.AfterFunc(time.Second*1, task)

//	//	select {
//	//	case <-timer1.C:
//	//		log.Println("executing task")

//	//	}

//}

//const INTERVAL_PERIOD time.Duration = 24 * time.Hour
const INTERVAL_PERIOD time.Duration = 1 * time.Second

const HOUR_TO_TICK int = 23
const MINUTE_TO_TICK int = 00
const SECOND_TO_TICK int = 03

type jobTicker struct {
	timer *time.Timer
}

func runningRoutine() {
	jobTicker := &jobTicker{}
	jobTicker.updateTimer()
	for {
		select {
		case <-jobTicker.timer.C:
			fmt.Println(time.Now(), "- just ticked")
			jobTicker.updateTimer()
		}
	}
}


func Print() {
	fmt.Println("Hello")
}


func ProcExpiration(f func()) {
	go func() {
		for {
			//f()
			now := time.Now()
			//fmt.Println("now",now)
			next := now.Add(time.Second * 5)
			//fmt.Println("next",next)
			next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(),next.Minute(),next.Second(), 0, next.Location())
			fmt.Println("next 2 ",next)
			//fmt.Println("next sub ",next.Sub(now))
			fmt.Println("----------------------")
			t := time.NewTimer(next.Sub(now))
			<-t.C
			f()
		}
	}()
}

func (t *jobTicker) updateTimer() {
	nextTick := time.Date(time.Now().Year(), time.Now().Month(),
		time.Now().Day(), HOUR_TO_TICK, MINUTE_TO_TICK, SECOND_TO_TICK, 0, time.Local)
	if !nextTick.After(time.Now()) {
		nextTick = nextTick.Add(INTERVAL_PERIOD)
	}
	fmt.Println(nextTick, "- next tick")
	diff := nextTick.Sub(time.Now())
	if t.timer == nil {
		t.timer = time.NewTimer(diff)
	} else {
		t.timer.Reset(diff)
	}
}

func findException() {
	for {
		log.Println("log")
		time.Sleep(time.Second * 3)
	}

}

func findExec() {
	c := cron.New()
	spec := "@every 1s"
	c.AddFunc(spec, task)
	c.Start()
	c.Run()
}

func main() {

	t := time.Date(2019, 8, 4, 23, 0, 0, 0, time.UTC)
	fmt.Println(t)

	//go findException()

	//findExec()
	//
	//// now add this signal to handle exit behavior, later, will add gracefully exit function
	//sigs := make(chan os.Signal, 1)
	//// receive notifications of the specified signals.
	//signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGSTOP)
	//<-sigs
	//fmt.Println("exit by ctrl-c or ctrl-z")

}
