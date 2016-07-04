package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type LogEntry struct {
	id        bson.ObjectId `json:"id" bson:"_id"`
	level     int           `json:"level" bson:"level"`
	message   string        `json:"message" bson:"message"`
	context   string        `json:"context" bson:"context"`
	timestamp time.Time     `json:"timestamp" bson:"timestamp"`
}

func (l LogEntry) Persist() error {
	session, err := GetMongoDbSession()
	defer session.Close()
	if err != nil {
		return err
	}
	session.DB("logger").C("logs").Insert(l)
	return nil
}
