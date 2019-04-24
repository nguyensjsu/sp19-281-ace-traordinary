/*
	picsqueryapi in Go (Version 1)
*/

package main

import 	"gopkg.in/mgo.v2/bson"


type Picture struct {
	PictureId    string   `json:"pictureId" bson:"pictureId"`
	UserId       string   `json:"userId" bson:"userId"`
	Title        string   `json:"title" bson:"title"`
	Price 	     float32  `json:"price" bson:"price"`
	Description  string   `json:"description" bson:"description"`
	TumbnailUrl  string   `json:"tumbnailUrl" bson:"tumbnailUrl"`
	OrigUrl      string   `json:"origUrl" bson:"origUrl"`
}	

type Payload struct {
	Id		     bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	PictureId    string   `json:"pictureId" bson:"pictureId"`
	UserId       string   `json:"userId" bson:"userId"`
	Title        string   `json:"title" bson:"title"`
	Price 	     float32  `json:"price" bson:"price"`
	Description  string   `json:"description" bson:"description"`
	TumbnailUrl  string   `json:"tumbnailUrl" bson:"tumbnailUrl"`
	OrigUrl      string   `json:"origUrl" bson:"origUrl"`
}

type Request struct {
	//QueryId       string  	`json:"queryId" bson:"queryId"`
	PictureId    string   `json:"pictureId" bson:"pictureId"`
	//UserId       string   `json:"userId" bson:"userId"`
	//PageNumber    string    `json:"pageNumber" bson:"pageNumber"`
	//PageSize         string    `json:"pageSize" bson:"pageSize"`
	//QueryStatus   string    `json:"queryStatus" bson:"queryStatus"`
}

type Log struct {
	QueryId       string  	`json:"queryId" bson:"queryId"`
	UserId        string  	`json:"userId" bson:"userId"`
	QueryStatus   string    `json:"queryStatus" bson:"queryStatus"`
}

var pictures map[string]Payload