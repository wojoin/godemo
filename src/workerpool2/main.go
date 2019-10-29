package main

import (
	"demo/src/workerpool2/task"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Test struct {
	Name string
}

func (t *Test) Exec() (interface{},error ){
	fmt.Println("Here is:", t.Name)
	msg := "Here is:"+t.Name
	return msg,nil
}

type Test2 struct {
	Name string
}

func (t *Test2) Exec() (interface{},error ){
	fmt.Println("Here is another:", t.Name)
	msg := "Here is another:" + t.Name
	return msg, nil
}



func entry2(c *gin.Context) {
	var work task.Job
	// fetch job
	t := Test{Name: "Coeus"}
	work = &t
	task.JobQueue <- work

	//t2 := Test2{Name: "Coeus"}
	//work = &t2
	//task.JobQueue <- work

	//fmt.Fprintf(res, "Hello World ...again")

	value, ok := <- task.Result
	if !ok {
		c.String(http.StatusBadRequest,"error")
	}

	c.String(http.StatusOK,"Result: %s",value)

}

func entry(res http.ResponseWriter, req *http.Request) {
	var work task.Job
	// fetch job
	t := Test{Name: "Coeus"}
	work = &t
	task.JobQueue <- work

	t2 := Test2{Name: "Coeus"}
	work = &t2
	task.JobQueue <- work

	//fmt.Fprintf(res, "Hello World ...again")
}

func main() {
	Port := "8086"
	IsHttp := true
	arg_num := len(os.Args)
	if 2 <= arg_num {
		Port = os.Args[1]
	}
	if 3 <= arg_num {
		if os.Args[2] == "true" {
			IsHttp = true
		} else {
			IsHttp = false
		}
	}
	fmt.Printf("server is http %t\n", IsHttp)
	fmt.Println("server listens at ", Port)

	router := gin.Default()

	apiWorker := router.Group("/api/workerpool")
	{
		apiWorker.GET("/work", entry2)
	}

	err := router.Run(":"+Port)
	if err != nil {
		panic("router run failed")
	}

	//http.HandleFunc("/", entry)

	//var err error
	//if IsHttp {
	//	err = http.ListenAndServe(":"+Port, nil)
	//} else {
	//	err = http.ListenAndServeTLS(":"+Port, "server.crt", "server.key", nil)
	//}
	//if err != nil {
	//	fmt.Println("Server failure /// ", err)
	//}
	//
	//fmt.Println("quit")
}