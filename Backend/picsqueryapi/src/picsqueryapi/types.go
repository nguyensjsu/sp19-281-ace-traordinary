/*
	picsqueryapi in Go (Version 1)
*/

package main

import "gopkg.in/mgo.v2/bson"

type Picture struct {
	ImageId     string  `json:"imageid" bson:"imageid"`
	UserId      string  `json:"userid" bson:"userid"`
	Title       string  `json:"title" bson:"title"`
	Price       float32 `json:"price" bson:"price"`
	Description string  `json:"description" bson:"description"`
	TumbnailUrl string  `json:"tumbnailUrl" bson:"tumbnailUrl"`
	OrigUrl     string  `json:"origurl" bson:"origurl"`
}

type Payload struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	ImageId     string        `json:"imageid" bson:"imageid"`
	UserId      string        `json:"userid" bson:"userid"`
	Title       string        `json:"title" bson:"title"`
	Price       float32       `json:"price" bson:"price"`
	Description string        `json:"description" bson:"description"`
	TumbnailUrl string        `json:"tumbnailUrl" bson:"tumbnailUrl"`
	OrigUrl     string        `json:"origurl" bson:"origurl"`
}

type Request struct {
	//QueryId       string  	`json:"queryId" bson:"queryId"`
	ImageId string `json:"imageid" bson:"imageid"`
	//UserId       string   `json:"userId" bson:"userId"`
	//PageNumber    string    `json:"pageNumber" bson:"pageNumber"`
	//PageSize         string    `json:"pageSize" bson:"pageSize"`
	//QueryStatus   string    `json:"queryStatus" bson:"queryStatus"`
}

type Log struct {
	QueryId     string `json:"queryId" bson:"queryId"`
	UserId      string `json:"userId" bson:"userId"`
	QueryStatus string `json:"queryStatus" bson:"queryStatus"`
}

var pictures map[string]Payload
