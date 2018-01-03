package models

import "gopkg.in/mgo.v2/bson"

type Case struct {
	Id 		bson.ObjectId `bson:"_id,omitempty"`
	Name 	string
}


