package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	//"github.com/jinzhu/configor"
)

type MessageProcess struct {
	Type int
}

func (m *MessageProcess) PrintMessage() {
	log.Println("MessageProcess", m.Type)
}

type Result struct {
	Rc  uint16
	Val interface{}
}

type Token struct {
	Ysid          string
	Ytid          string
	UserId        uint
	Ts            string
	Authorization string
}

const (
	GPSTopic      = "$share/group1/gps/#"
	StatusTopic   = "$share/group2/status/#"
	EventTopic    = "$share/group3/event/#"
	LocationTopic = "$share/group4/location/#"

	GpsCol    = "gps"
	StatusCol = "status"
	EventCol  = "event"
)

var TopicList = []string{GPSTopic, StatusTopic, EventTopic, LocationTopic}

type MsgHeader struct {
	Version   string `json:"version,omitempty"`
	Type      int    `json:"type,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

type MsgStatus struct {
	Carstate   int     `json:"car_state,omitempty"`
	Carerror   string  `json:"car_error,omitempty"`
	Carstage   int     `json:"car_stage,omitempty"`
	Carspeed   float64 `json:"car_speed,omitempty"`
	Odom       int     `json:"odom,omitempty"`
	Perception int     `json:"perception,omitempty"`
	Planner    int     `json:"planner,omitempty"`
}

type MqttMessage struct {
	Version   string `json:"version,omitempty"`
	Type      int    `json:"type,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`

	MessageData MsgStatus `json:"data,omitempty"`
}

type CommandType int

const (
	_ CommandType = iota
	CALLC
	TRANSC
	RETURNC
	END
)

func IsValid(cmd int) bool {
	return cmd > 0 && cmd < int(END)
}

type Location struct {
	X int
	Y int
}

type Point struct {
	X int
	Y int
	Location
}

type Color struct {
	C int
}

type ColorPoint struct {
	Point
	Color
}

func (l *Location) Scale(factor int) *Location {
	fmt.Println("promotion twice, method Scale in Location")
	return &Location{l.X * factor, l.Y * factor}
}

func (p *Point) Scale(factor int) *Point {
	fmt.Println("promotion once, method Scale in Point")
	return &Point{(p.X) * factor, (p.Y) * factor, Location{p.X * factor, p.Y * factor}}
}

func (p *Color) Scale(factor int) *Color {
	fmt.Println("promotion once, method Scale in Color")
	return &Color{p.C * factor}
}

func (p *ColorPoint) Scale(factor int) *ColorPoint {
	fmt.Println("directly, method Scale in ColorPoint")
	return &ColorPoint{Point{p.X * factor, p.Y * factor, Location{p.X * factor, p.Y * factor}}, Color{p.C * factor}}
}

func getRole(userroles []string) int {
	var maxRole = 0
	for _, r := range userroles {
		var currole = 0
		if strings.Contains(r, "guest") {
			currole = 0
		} else if strings.Contains(r, "developer") {
			currole = 1
		} else if strings.Contains(r, "admin") || strings.Contains(r, "owner") {
			currole = 2
		} else if strings.Contains(r, "system") {
			currole = 3
		}

		if maxRole < currole {
			maxRole = currole
		}
	}

	return maxRole
}


func RemoveRepeat(a []string, b []string) []string {
	res := []string{}
	for _, av := range a {
		repflag := false
		for _, resv := range res {
			if resv == av {
				repflag = true
				break
			}
		}
		if !repflag {
			res = append(res, av)
		}
	}
	for _, av := range b {
		repflag := false
		for _, resv := range res {
			if resv == av {
				repflag = true
				break
			}
		}
		if !repflag {
			res = append(res, av)
		}
	}
	return res
}

func RemoveRepeat2(a []string, b []string) []string {
	tmpMap := make(map[string]struct{})
	for _, av := range a {
		tmpMap[av] = struct{}{}
	}
	for _, bv := range b {
		tmpMap[bv] = struct{}{}
	}
	res := []string{}
	for key := range tmpMap {
		res = append(res, key)
	}
	return res
}

type UserEntity struct {
	UserID     int64
	UROle      string
	Expiration string
}

//type Slice []interface{}
//
//func removeDup(data Slice) (result Slice) {
//	resultMap := map[string]bool{}
//	for _, v := range data {
//		data, _ := json.Marshal(v)
//		resultMap[string(data)] = true
//	}
//
//	for k := range resultMap {
//		var t interface{}
//		if err := json.Unmarshal([]byte(k), &t); err != nil {
//			log.Printf("removeDuplicate, origin data: %+v error: %s", data, err)
//			return nil
//		}
//		result = append(result, t)
//	}
//	return result
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
		removeDup(v.Index(0))
	case reflect.Struct:
		dataType := reflect.New(v.Type())
		log.Printf("type: %s",dataType)

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


func removeDuplicate(v reflect.Value) interface{} {
	resultMap := map[string]bool{}
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Print("%s = invalid\n")
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


func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func display(path string, v reflect.Value) {
	resultMap := map[string]bool{}

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			str := fmt.Sprintf("%+v",v.Index(i))
			data, _ := json.Marshal(str)
			resultMap[string(data)] = true
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}

		resultMap[v.String()] = true
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}


type Config struct {
	Name string
	Meta struct {
		Desc string
		Properties map[string]string
		Users []string
	}
}

func initializeStruct(t reflect.Type, v reflect.Value) {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		ft := t.Field(i)
		switch ft.Type.Kind() {
		case reflect.Map:
			f.Set(reflect.MakeMap(ft.Type))
		case reflect.Slice:
			f.Set(reflect.MakeSlice(ft.Type, 0, 0))
		case reflect.Chan:
			f.Set(reflect.MakeChan(ft.Type, 0))
		case reflect.Struct:
			initializeStruct(ft.Type, f)
		case reflect.Ptr:
			fv := reflect.New(ft.Type.Elem())
			initializeStruct(ft.Type.Elem(), fv.Elem())
			f.Set(fv)
		default:
		}
	}
}

//func RemoveRepeat(a []string, b []string) []string {
//	tmpMap := make(map[string]struct{})
//	for _, av := range a {
//		tmpMap[av] = struct{}{}
//	}
//	for _, bv := range b {
//		tmpMap[bv] = struct{}{}
//	}
//	res := []string{}
//	for key := range tmpMap {
//		res = append(res, key)
//	}
//	return res
//}

func RemoveDup(data []string) []string {
	resultMap := map[string]bool{}
	for _, v := range data {
		data, _ := json.Marshal(v)
		resultMap[string(data)] = true
	}
	result := []string{}
	for k := range resultMap {
		var t string
		if err := json.Unmarshal([]byte(k), &t); err != nil {
			return nil
		}
		result = append(result, t)
	}
	return result
}

func CronTask() {
	fmt.Println("cron task start...")
	time.Sleep(time.Second * 4)
	fmt.Println("cron task finished")
}


func main() {
	t := reflect.TypeOf(Config{})
	v := reflect.New(t)
	initializeStruct(t, v.Elem())
	c := v.Interface().(*Config)
	c.Meta.Properties["color"] = "red" // map was already made!
	c.Meta.Users = append(c.Meta.Users, "srid") // so was the slice.
	fmt.Println(v.Interface())


	//dupData := Slice{"a","b","c","a"}
	//log.Printf("result: %+v",removeDuplicate(dupData))

	//sliceData := []UserEntity{
	//	{UserID:123456,UROle:"developer",Expiration:"2019-06-26"},
	//	{UserID:123456,UROle:"developer",Expiration:"2019-06-27"},
	//	{UserID:123456,UROle:"developer",Expiration:"2019-06-26"},
	//	{UserID:123456,UROle:"maintainer",Expiration:"2019-06-26"}}

	// elem is array
	//dupStructData := Slice{sliceData}

	//log.Printf("dup struct result: %+v",removeDuplicate(dupStructData))

	dupStructData2 := []UserEntity{
		UserEntity{UserID:123456,UROle:"developer",Expiration:"2019-06-26"},
		UserEntity{UserID:123456,UROle:"developer",Expiration:"2019-06-27"},
		UserEntity{UserID:123456,UROle:"developer",Expiration:"2019-06-26"},
		UserEntity{UserID:123456,UROle:"maintainer",Expiration:"2019-06-26"}}

	//log.Printf("dup struct result2: %+v",removeDuplicate(reflect.ValueOf(dupStructData2)))
	res := RemoveDuplicate(dupStructData2)
	log.Printf("dup struct result2: %+v",res)

	//Display("dupStructData2", dupStructData2)

//	elemType := reflect.TypeOf(sliceData)
//	log.Printf("type: ", elemType)
//	log.Printf("type kind: ", elemType.Kind().String())
//
//
//	t1 := time.Now()
//
//	t2 := time.Now().Add(10 * time.Second)
//	timeSpan := int(t2.Sub(t1).Seconds())
//	log.Println("time span: ", timeSpan)
//
//	if timeSpan > 5 {
//		log.Println("expired")
//	}
//
//	messageProc := &MessageProcess{Type: 1}
//	messageProc.PrintMessage()
//
//	var carcmd CommandType = RETURNC
//	fmt.Println("----carcmd", carcmd)
//	fmt.Println(RETURNC)
//	fmt.Println(TRANSC)
//
//	cmd := 2
//	var transCmd = CommandType(cmd)
//
//	log.Println("---------trans cmd", transCmd)
//
//	var increment uint64
//
//	log.Println("increment------", increment)
//	log.Println("increment + 1", atomic.AddUint64(&increment, 1))
//
//	var f float32
//	var i int32
//
//	fmt.Println("float32 size: ", unsafe.Sizeof(f))
//	fmt.Println("int64 size: ", unsafe.Sizeof(i))
//
//	fmt.Println(IsValid(1))
//
//	fmt.Println(time.Now().Format(time.RFC3339))
//
//	var t time.Time
//
//	log.Println("----------time", t)
//
//	if t.IsZero() {
//		log.Println("---------zero time")
//	}
//
//	for index, topic := range TopicList {
//		fmt.Println(index, "--", topic)
//	}
//
//	data := []byte(`{
//    "version": "v1",
//    "type": 1,
//    "timestamp": 1541390405000,
//    "data": {
//        "car_state": 1,
//        "car_error": "some error description",
//        "car_stage": 1,
//        "car_speed": 1.6,
//        "odom": 50,
//        "perception": 10,
//        "planner": 10
//    }
//}`)
//
//	var mqttmsg MqttMessage
//
//	if err := json.Unmarshal(data, &mqttmsg); err != nil {
//		fmt.Println("json unmarshal error %s", err.Error())
//		return
//	}
//
//	fmt.Println(mqttmsg)
//
//	nums := []int{2, 3, 4}
//	sum := 0
//	for i, num := range nums {
//		fmt.Println("i = ", i)
//		sum += num
//	}
//
//	var iii int = 1
//	log.Printf("integer %q", iii)
//
//	// select statement for multiplexing
//
//	log.Println("rocket launch")
//	abort := make(chan struct{})
//	select {
//	case <-time.After(2 * time.Second):
//		log.Println("time flies")
//		log.Println("rocket has launched")
//	case <-abort:
//		log.Println("abort")
//		return
//	}
//
//	bytedata := []byte{0x2, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x73, 0x8, 0x2, 0x0, 0x0, 0x1}
//	var a int
//	err := binary.Read(bytes.NewReader(bytedata[0:4]), binary.LittleEndian, &a)
//	if err != nil {
//		fmt.Println("read error: ", err.Error())
//	}
//	r := int(a)
//	fmt.Println("read data: ", r)
//	fmt.Println(bytedata[0:4])
//
//	var mySlice = []byte{0010, 0010}
//	data2 := binary.LittleEndian.Uint16(mySlice)
//	fmt.Println(data2)
//
//	fmt.Println("--------------------Embedded type--------------")
//	l := Location{1, 1}
//	p := Point{1, 1, l}
//	//pp := p.Scale(2)
//	//fmt.Println(*pp)
//
//	cp := ColorPoint{p, Color{1}}
//	fmt.Println(cp.Scale(3))
//
//
//	log.Println("maxRole: ", getRole([]string{"guest", "developer", "system"}))
//
//	type User struct{
//		UserID string
//		Age int
//	}
//
//	u := User{UserID: "123", Age: 28}
//	log.Printf("user info : %+v", u)
//
//	stra := []string{"create", "delete", "view","create"}
//	//strb := []string{"delete", "setting"}
//
//	var b []string
//	resDup := RemoveRepeat2(stra, b)
//	log.Printf("result: %v",resDup)
//	log.Printf("time: %s",time.Now().Format("2006-01-02 15:04:05"))

	src := []string{"a","b","c","a"}
	src = RemoveDup(src)
	fmt.Println("remove deplicate: ",src)

	timeLocal := time.FixedZone("CST",3600*8)
	time.Local = timeLocal
	now := time.Now().Local()

	fmt.Println("time now: %s",now)

	for{
		now := time.Now()
		//next := now.Add(time.Hour * 5)
		//next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location()) //获取下一个凌晨的日期
		next := now.Add(time.Hour * 24)
		log.Println("next:",next)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		log.Println("after duration: ",next.Sub(now))
		t := time.NewTimer(next.Sub(now))
		<-t.C

		go CronTask()
	}
}
