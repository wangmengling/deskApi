package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionArticle holds the name of the articles collection
	CollectionColors = "colors"
)

type Colors struct {
	//Id 		bson.ObjectId `bson:"_id,omitempty"`
	Id bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name 	string  `bson:"name"`
}
