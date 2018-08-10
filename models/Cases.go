package models

import (
	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionArticle holds the name of the articles collection
	CollectionCases = "cases"
)

type Cases struct {
	//Id 		bson.ObjectId `bson:"_id,omitempty"`
	Id bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title 	string  `bson:"title"`
	Time    string  `bson:"time"`
	Username string  `bson:"username"`
	Contact string  `bson:"contact"`
	Phone string  `bson:"phone"`
	Style string  `bson:"style"`
	Color string  `bson:"color"`
	Address string  `bson:"address"`
	CreateTime string  `bson:"createTime"`
	ThumbUrl string  `bson:"thumbUrl"`
	VideoUrl string  `bson:"videoUrl"`
	Description string `bson:"description"`
	ImageUrl [][]bson.M  `bson:"imageUrl"`
}
