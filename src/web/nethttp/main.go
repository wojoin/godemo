package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type LoginInfo struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
}

func index(w http.ResponseWriter, r *http.Request) {
	var login LoginInfo

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Println(login)
	fmt.Println("login")
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, "register device success ")
	//fmt.Fprintf(w, login.Name)
	//fmt.Fprintf(w, login.Pwd)
}

func main() {
	http.HandleFunc("/registry/device", index)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
