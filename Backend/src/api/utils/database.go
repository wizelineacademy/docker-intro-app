package utils

import (
	"errors"
	"fmt"
	"sync"
	"os"

	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

var DBSession *mgo.Session
var once sync.Once

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getConnection() (db *mgo.Session, err error) {
	once.Do(func() {

		session, errordb := mgo.Dial(getenv("MONGODB_PORT_27017_TCP_ADDR", "127.0.0.1") + ":" + getenv("MONGODB_PORT_27017_TCP_PORT", "27017"))
		if err != nil {
			panic(err)
		}
		session.SetMode(mgo.Monotonic, true)
		// Collection People
		err = errordb
		DBSession = session
		return
	})

	return DBSession, err
}

func GetDB() (*mgo.Session, error) {
	db, err := getConnection()
	//db.LogMode(CONFIG.DebugMode)
	if err != nil {
		fmt.Println("error connecting to db", err.Error())
		return nil, errors.New("cannot connect to db")
	}
	return db, nil
}
