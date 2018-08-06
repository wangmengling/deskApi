package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
<<<<<<< HEAD

=======
	"os"
>>>>>>> 5f0d6c036d6f1eb9b3b158437e7fccb8f3be2c8e
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

const (
	// MongoDBUrl is the default mongodb url that will be used to connect to the
	// database.
	MongoDBUrl = "mongodb://localhost:27017/sonnet"
)

// Connect connects to mongodb
func Connect() {
	uri := os.Getenv("MONGODB_URL")
	//
	if len(uri) == 0 {
		uri = MongoDBUrl
	}

	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}

//func GetMgo() *mgo.Session {
//	return Session
//}
//
//func GetDataBase() *mgo.Database {
//	return database
//}
//
//func GetErrNotFound() error {
//	return mgo.ErrNotFound
//}