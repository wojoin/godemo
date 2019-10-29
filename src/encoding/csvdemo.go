package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Detail struct {
	Reason []string `json:"reason"`
	Msg []string `json:"Msg"`
}

type Result struct {
	Status       int    `json:"status"`
	Message      string `json:"message"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Detail `json:"detail"`
}


func main() {
	in := `first_name;last_name;username
"Rob";"Pike";rob
# lines beginning with a # character are ignored
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
	r := csv.NewReader(strings.NewReader(in))
	r.Comma = ';'
	r.Comment = '#'

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(records)
	fmt.Println("------json")

	json_str0 := `{"status":0, "message":"success"}`
	json_str1 := `{"status":1, "error_code":5, "error_message":"error"}`

	res0 := Result{}
	res1 := Result{}

	err0 := json.Unmarshal([]byte(json_str0), &res0)
	err1 := json.Unmarshal([]byte(json_str1), &res1)

	fmt.Println(res0, err0)
	fmt.Println(res1, err1)
}
