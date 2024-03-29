package config

import (
	"os"

	mgo "gopkg.in/mgo.v2"
)

func GetMongoDB() (*mgo.Database, error) {
	host := os.Getenv("MONGO_HOST")
	dbname := os.Getenv("MONGO_DB_NAME")

	session, err := mgo.Dial(host)

	if err != nil {
		return nil, err
	}

	db := session.DB(dbname)

	return db, nil
}
