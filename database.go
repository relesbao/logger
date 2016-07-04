package main

import (
	"gopkg.in/mgo.v2"
)

const mongoDbHostname string = "127.0.0.1"

func GetMongoDbSession() (*mgo.Session, error) {
	return mgo.Dial(mongoDbHostname)
}
