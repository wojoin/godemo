package db

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"os"
	"demo/src/casbin/conf"
)

var (
	// Session stores mongo session
	Session *mgo.Session
	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

// Connect connects to mongodb
func Connect(mongoDBUrl string) {
	uri := os.Getenv("MONGODB_URL")
	if len(uri) == 0 {
		uri = mongoDBUrl
	}

	mongo := &mgo.DialInfo{
		Addrs: []string{conf.AppConfig.MONGO.URL},
		//Source:   conf.AppConfig.MONGO.DBNAME,
		Username: conf.AppConfig.MONGO.USERNAME,
		Password: conf.AppConfig.MONGO.PASSWORD,
		Database: conf.AppConfig.MONGO.DBNAME,
	}

	//mongo, err := mgo.ParseURL(uri)
	log.Infof("mongodb uri: %+v", uri)

	//s, err := mgo.DialWithInfo(mongo)
	s, err := mgo.Dial(uri)
	if err != nil {
		log.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	//s.SetSafe(&mgo.Safe{})
	s.EnsureSafe(&mgo.Safe{W: 1, WMode: "majority", FSync: true})
	s.SetMode(mgo.Eventual, true)
	//log.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}