package task

import "runtime"

var (
	MaxWorker = runtime.NumCPU()
	MaxQueue  = 512
)

type Job interface {
	Exec() (interface{},error)
}

var JobQueue chan Job

var Result = make(chan interface{},MaxQueue)
