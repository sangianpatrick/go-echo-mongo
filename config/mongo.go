package config

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

// GetMongoDB is a function that will return MongoDB instance
func GetMongoDB() (*mgo.Database, error) {
	host := "localhost:27017"
	dbName := "mongotest"
	user := "mongotest"
	pwd := "14qwafzx"

	dialInfo, err := mgo.ParseURL(host)
	if err != nil {
		return nil, err
	}
	dialInfo.Timeout = 5 * time.Second
	dialInfo.Username = user
	dialInfo.Password = pwd
	dialInfo.Database = dbName
	dialInfo.Mechanism = ""

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}

	db := session.DB(dbName)
	return db, nil
}
