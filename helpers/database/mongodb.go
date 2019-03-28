package database

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

// GetMongoDB is a function that will return MongoDB instance
func GetMongoDB(c map[string]string) (*mgo.Database, error) {
	dialInfo, err := mgo.ParseURL(c["host"])
	if err != nil {
		return nil, err
	}
	dialInfo.Timeout = 5 * time.Second
	dialInfo.Username = c["user"]
	dialInfo.Password = c["password"]
	dialInfo.Database = c["db"]
	dialInfo.Mechanism = ""

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, err
	}

	db := session.DB(c["db"])
	return db, nil
}
