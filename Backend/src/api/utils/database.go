package utils

import (
	"errors"
	"fmt"
	"sync"

	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

var DBSession *mgo.Session
var once sync.Once

func getConnection() (db *mgo.Session, err error) {
	once.Do(func() {
		session, errordb := mgo.Dial("127.0.0.1:27017")
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