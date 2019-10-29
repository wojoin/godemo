package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

type Person struct {
	ID   int    `tag_name:"tag 1"`
	Name string `tag_name:"tag 2"`
}

var person = Person{
	ID:   1,
	Name: "join",
}

type People []Person

var pepple = People{
	Person{ID:   2, Name: "2"},
	Person{ID:   3, Name: "3"},
	Person{ID:   4, Name: "4"}}

type UserEntity struct {
	UserID     int64  `json:"uid,omitempty"`
	UROle      string `json:"urole,omitempty"`
	Expiration string `json:"expiration,omitempty"`
}

//func RemoveDuplicate(x interface{}) interface{} {
//	return removeDuplicate(reflect.TypeOf(x), reflect.ValueOf(x))
//}

func RemoveDuplicate(x interface{}) interface{} {
	log.Printf("remove duplicate, value: %+v", x)
	return removeDup(reflect.ValueOf(x))
}

func removeDup(v reflect.Value) interface{} {
	resultMap := map[string]bool{}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Println("invalid type")
		return nil
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			str := fmt.Sprintf("%+v", v.Index(i))
			data, _ := json.Marshal(str)
			resultMap[string(data)] = true
		}
	}

	result := []interface{}{}
	for k := range resultMap {
		var t interface{}
		if err := json.Unmarshal([]byte(k), &t); err != nil {
			return nil
		}
		result = append(result, t)
	}
	return result
}


func removeDuplicate(x interface{}) interface{} {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	resultMap := map[string]bool{}
	if t.Kind() != reflect.Slice {
		fmt.Println("The interface is not a slice.")
		return nil
	}



	for i := 0; i < v.Len(); i++ {
		str := fmt.Sprintf("%+v", v.Index(i)) // TODO how to get struct
		data, _ := json.Marshal(str) // should marshal struct
		//data, _ := json.Marshal(v.Index(i))
		resultMap[string(data)] = true
	}


	newLen := v.Len()
	newCap := newLen
	elemType := t.Elem()
	//log.Println("elemType: ", elemType.String())

	res := reflect.MakeSlice(reflect.SliceOf(elemType), newLen, newCap)


	var slice interface{}
	//result := []interface{}{}
	for k := range resultMap {
		var t interface{}
		if err := json.Unmarshal([]byte(k), &t); err != nil {
			return nil
		}
		// TODO
		//  1. recover struct like "{UserID:123456 UROle:developer Expiration:2019-06-27}"
		//  2. assign to entity and
		//  3. then append to slice res
		log.Println("recover struct from json.Unmarshal: ",k)

		//result = append(result, t)
		//res = reflect.Append(res, t)
	}

	slice = reflect.ValueOf(res).Interface()
	return slice
}

//func removeDup(t reflect.Type, v reflect.Value) interface{} {
//	resultMap := map[string]bool{}
//	log.Println("type removeDup:",t)
//	tmp := t.Kind()
//	log.Println(tmp)
//
//		for i := 0; i < v.Len(); i++ {
//			str := fmt.Sprintf("%+v", v.Index(i))
//			data, _ := json.Marshal(str)
//			resultMap[string(data)] = true
//		}
//
//		var t interface{}
//		newLen := v.Len()
//		newCap := newLen
//		elemType := t.
//
//		removeDup(reflect.TypeOf(v), reflect.ValueOf(v))
//
//
//	result := []interface{}{}
//	for k := range resultMap {
//		var t interface{}
//		if err := json.Unmarshal([]byte(k), &t); err != nil {
//			return nil
//		}
//		result = append(result, t)
//	}
//	return result
//}


func main() {
	log.Print(reflect.TypeOf(2))

	v := reflect.ValueOf(person)
	//v2 := reflect.ValueOf(pepple).Elem()
	fmt.Println(v)

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		fmt.Printf("Field Value: %v\r\n", f)
	}

	fmt.Println(reflect.ValueOf(0).Type().Kind())

	dupStructData2 := []UserEntity{
		UserEntity{UserID:123456,UROle:"developer",Expiration:"2019-06-26"},
		UserEntity{UserID:123456,UROle:"developer",Expiration:"2019-06-27"},
		UserEntity{UserID:123456,UROle:"developer",Expiration:"2019-06-26"},
		UserEntity{UserID:123456,UROle:"maintainer",Expiration:"2019-06-26"}}

	log.Printf("struct value: %+v",dupStructData2)
	res := removeDuplicate(dupStructData2)
	log.Println(res)
}
