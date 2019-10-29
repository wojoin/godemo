package main

import (
	"demo/src/conf"
	"fmt"
	"github.com/jinzhu/configor"
	"net/url"
	"os"
	"strconv"

	//"io/ioutil"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


var (
	// Session stores mongo session
	Session *mgo.Session
	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)


type Person struct {
	Id         bson.ObjectId `bson:"_id"`
	EmployeeID string        `bson:"employeeid"`
	Name       string        `bson:"name"`
	Phone      string        `bson:"phone"`
}

type Map struct {
	Id          bson.ObjectId `bson:"_id"`
	Map_id      string        `bson:"map_id"`
	Description string        `json:"desc" bson:"description"`
	Content     string        `bson:"content"`
}

const (
	URL        = "localhost:27017"
	DBNAME     = "horizon"
	COLLPERSON = "person"
	COLLMAP    = "map"
)

var (
	mgoSession *mgo.Session
)

func getSession() *mgo.Session {
	fmt.Println("get session")
	if mgoSession == nil {
		var err error
		maxWait := time.Duration(5 * time.Second)

		mgoSession, err = mgo.DialWithTimeout(URL, maxWait)
		if err != nil {
			panic(err)
		}
	}
	return mgoSession.Clone()
}

// get collection object and then operation with it.
func withCollection(collection string, op func(*mgo.Collection) error) error {
	session := getSession()
	defer session.Close()
	c := session.DB(DBNAME).C(collection)
	return op(c)
}

// add persion collection
func AddPersion(p *Person) bool {
	p.Id = bson.NewObjectId()
	insertion := func(c *mgo.Collection) error {
		return c.Insert(p)
	}
	err := withCollection(COLLPERSON, insertion)
	if err != nil {
		return false
	}
	return true
}

func GetPersonByID(id string) *Person {
	objid := bson.ObjectIdHex(id)
	p := new(Person)

	query := func(c *mgo.Collection) error {
		return c.FindId(objid).One(&p)
	}
	withCollection(COLLPERSON, query)
	return p
}

func GetPersons() []Person {
	var persons []Person
	query := func(c *mgo.Collection) error {
		return c.Find(nil).All(&persons)
	}
	err := withCollection(COLLPERSON, query)
	if err != nil {
		return persons
	}
	return persons
}

func UpdatePerson(query bson.M, change bson.M) bool {
	update := func(c *mgo.Collection) error {
		return c.Update(query, change)
	}
	err := withCollection(COLLPERSON, update)
	if err != nil {
		log.Println("err--------", err.Error())
		return false
	}
	return true
}

/**
 * 执行查询，此方法可拆分做为公共方法
 * [SearchPerson description]
 * @param {[type]} collectionName string [description]
 * @param {[type]} query          bson.M [description]
 * @param {[type]} sort           bson.M [description]
 * @param {[type]} fields         bson.M [description]
 * @param {[type]} skip           int    [description]
 * @param {[type]} limit          int)   (results      []interface{}, err error [description]
 */
func SearchPerson(collectionName string, query bson.M, sort string, fields bson.M, skip int, limit int) (results []interface{}, err error) {
	search := func(c *mgo.Collection) error {
		return c.Find(query).Sort(sort).Select(fields).Skip(skip).Limit(limit).All(&results)
	}
	err = withCollection(collectionName, search)
	return
}

func GetMapByID(mapid string) *Map {
	result := &Map{}
	session := getSession()
	defer session.Close()
	coll := session.DB(DBNAME).C(COLLMAP)
	err := coll.Find(bson.M{"map_id": mapid}).One(&result)
	if err != nil {
		log.Println("mongodb find error")
		return result
	}

	return result

}

func mongoConnect(mongoUri string) {
	uri := os.Getenv("MONGODB_URL")
	if len(uri) == 0 {
		uri = mongoUri
	}
	mongo := &mgo.DialInfo{
		Addrs: []string{conf.AppConfig.MONGO.URL},
		//Source:   conf.AppConfig.MONGO.DBNAME,
		Username: conf.AppConfig.MONGO.USERNAME,
		Password: conf.AppConfig.MONGO.PASSWORD,
		Database: conf.AppConfig.MONGO.DBNAME,
	}
	//mongo, err := mgo.ParseURL(uri)
	log.Println("mongo uri: %+v",uri)

	//s, err := mgo.DialWithInfo(mongo)
	s, err := mgo.Dial(uri)
	if err != nil {
		log.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})

	//dialInfo := mongo
	//session := s
	//log.Println("dialInfo: %+v",dialInfo)
	//log.Println("session: %+v",session)

	Mongo = mongo
	Session = s
}

type UserDB struct {
	Uid   string   `json:"uid" bson:",omitempty"`   //用户ID
	UName string   `json:"uname" bson:",omitempty"` //用户名，LDAP的邮箱前缀，不可重复
	Dids  []string `json:"dids" bson:",omitempty"`  // 用户所属的部门ID
	Ctime string `json:"ctime" bson:",omitempty"` //创建时间
	Mtime string `json:"mtime" bson:",omitempty"` //修改时间
}

func GetUserByID(uid string) (*UserDB, error) {
	s := Session.Clone()
	defer s.Close()
	db := s.DB(Mongo.Database).C("toolchain_users")
	user := &UserDB{}
	err := db.Find(bson.M{"uname": uid}).One(&user)
	if err != nil {
		return nil, err
	}
	log.Printf("get users array: %v", user)
	return user, nil
}

func main() {
	//	p := Person{
	//		EmployeeID: "42",
	//		Name:       "join",
	//		Phone:      "18688888888",
	//	}
	//	AddPersion(&p)

	//	content, err := ioutil.ReadFile("/home/join/zgc.xodr")
	//	if err != nil {
	//		log.Fatal(err)
	//	}

	//	fmt.Printf("File contents: %s", content)

	//	//xmlcontent := "read from file"

	//	m := Map{
	//		Map_id:      "123",
	//		Description: "shanghai map",
	//		Content:     string(content),
	//	}

	//	AddMap(&m)

	if err := configor.New(&configor.Config{Debug:false}).Load(&conf.AppConfig,"src/conf/config.yaml"); err != nil {
		log.Println("load config error: %s", err.Error())
		return
	}

	//m := GetMapByID("345")
	//if m != nil {
	//
	//}

	//name := "join1999"
	//selector := bson.M{"employeeid": "42"}

	//data := bson.M{"$set": bson.M{"name": "join1999"}}

	//flag := UpdatePerson(bson.M{"employeeid": "42"}, bson.M{"$set": bson.M{"name": "join1999"}})
	//if flag {
	//	log.Println("update success")
	//}

	mongodbURL := ""
	//if conf.AppConfig.MONGO.USERNAME != "" && conf.AppConfig.MONGO.PASSWORD!= "" {
	//	mongodbURL = "mongodb://" + url.PathEscape(conf.AppConfig.MONGO.USERNAME) + ":" +
	//		url.PathEscape(conf.AppConfig.MONGO.PASSWORD) + "@" +
	//		conf.AppConfig.MONGO.URL + "/" + url.PathEscape(conf.AppConfig.MONGO.DBNAME)
	//} else {
	//	mongodbURL = "mongodb://" + conf.AppConfig.MONGO.URL + "/" + url.PathEscape(conf.AppConfig.MONGO.DBNAME)
	//}

	if conf.AppConfig.MONGO.USERNAME != "" && conf.AppConfig.MONGO.PASSWORD!= "" {
		mongodbURL = "mongodb://" + url.PathEscape(conf.AppConfig.MONGO.USERNAME) + ":" +
			//url.PathEscape(conf.AppConfig.MONGO.PASSWORD) + "@" + conf.AppConfig.MONGO.URL + "/" + conf.AppConfig.MONGO.DBNAME
			url.PathEscape(conf.AppConfig.MONGO.PASSWORD) + "@" + conf.AppConfig.MONGO.URL
	} else {
		mongodbURL = "mongodb://" + conf.AppConfig.MONGO.URL + "/" + url.PathEscape(conf.AppConfig.MONGO.DBNAME)
	}

	log.Println(mongodbURL)
	mongoConnect(mongodbURL)

	res, _ := GetUserByID("605117124")
	if res == nil {
		return
	}


	restype := "project"
	role := 3

	log.Println(restype+strconv.Itoa(role))


}
