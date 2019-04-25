package main

import "time"

type Reaction struct{
	Image_id 		string   	`bson:"imageId,omitempty"`
	Reaction_type 	string		`json:"ReactionType,omitempty" bson:"reactionType,omitempty"`
	UserId			string 		`json:"UserId,omitempty" bson:"userId,omitempty"`
	Username 		string 	 	`json:"Username,omitempty" bson:"userName,omitempty"`
	Timestamp 		time.Time	`json:"Timestamp,omitempty" bson:"timeStamp,omitempty"`
	Comment 		string  	`json:"Comment,omitempty" bson:"comment,omitempty"`
}

type ImageReaction struct {
	Imageid 	string		`json:"ImageId,omitempty"`
	likes	 	[]Likes		`json:",omitempty"`
	comments	[]Comments	`json:"Comments,omitempty"`

}

type Comments struct{
	Username	string 		`json:"Username,omitempty"bson:"userName,omitempty"`
	Comment		string 		`json:"Comment,omitempty"bson:"comment,omitempty"`
	TimeStamp 	string 	`json:"TimeStamp,omitempty"bson:"timeStamp,omitempty"`
}

type Likes struct{
	Username		string 			`json:"Username,omitempty"bson:"userName,omitempty"`
	TimeStamp 		time.Time 		`json:"TimeStamp,omitempty"bson:"timeStamp,omitempty"`
}