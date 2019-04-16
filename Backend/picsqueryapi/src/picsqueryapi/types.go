/*
	picsqueryapi in Go (Version 1)
*/

package main


type Picture struct {
	PictureId    string   `json:"pictureId"`
	PictureTitle string   `json:"pictureTitle"`
	Price 	     float32  `json:"price"`
	Description  string   `json:"description"`
	TumbnailUrl  string   `json:"tumbnailUrl"`
	OrigUrl      string   `json:"origUrl"`
}	
type PictureQuery struct {
	QueryId     string  `json:"queryId" bson:"queryId"`
	UserId      string  `json:"userId" bson:"userId"`
	PictureId   string  `json:"pictureId" bson:"pictureId"`
	PageNumber  string  `json:"pageNumber" bson:"pageNumber"`
	Count       string  `json:"count" bson:"count"`
	QueryStatus string  `json:"queryStatus" bson:"queryStatus"`
	IpAddress	string	`json:"ipaddress" bson:"ipaddress"`
}

type Payload struct {
	QueryId       string  	`json:"queryId" bson:"queryId"`
	UserId        string  	`json:"userId" bson:"userId"`
	PictureId     string 	`json:"pictureId"`
	PictureTitle  string 	`json:"pictureTitle"`
	Price 	      float32	`json:"price"`
	Description   string    `json:"description"`
	TumbnailUrl   string    `json:"tumbnailUrl"`
	OrigUrl       string    `json:"origUrl"`
	PageNumber    string    `json:"pageNumber" bson:"pageNumber"`
	Count         string    `json:"count" bson:"count"`
	QueryStatus   string    `json:"queryStatus" bson:"queryStatus"`
}

var queries map[string]PictureQuery