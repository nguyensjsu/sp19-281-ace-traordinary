package main

import "time"

type Reaction struct{
	Image_id 		string   	`bson:"imageId,omitempty"`
	Reaction_type 	string		`json:"ReactionType,omitempty" bson:"reactionType,omitempty"`
	UserId			string 		`json:"UserId,omitempty" bson:"userId,omitempty"`
	Username 		string 	 	`json:"Username,omitempty" bson:"userName,omitempty"`
	Timestamp 		time.Time	`json:"Timestamp,omitempty" bson:"timeStamp,omitempty"`
	Comment 		string  	`json:"Comment,omitempty" bson:"comment,omitempty"`
	CommentId		string		`json:"CommentId,omitempty"bson:"commentId,omitempty"`
}

type ImageReaction struct {
	Imageid 	string		`json:"ImageId,omitempty"`
	Likes	 	[]Likes		`json:",omitempty"`
	Comments	[]Comments	`json:"Comments,omitempty"`

}

type Comments struct{
	//Imageid		string		`json:"Imageid,omitempty"bson:"imageId,omitempty"`
	Username	string 		`json:"Username,omitempty"bson:"userName,omitempty"`
	Comment		string 		`json:"Comment,omitempty"bson:"comment,omitempty"`
	TimeStamp 	string 		`json:"TimeStamp,omitempty"bson:"timeStamp,omitempty"`
	CommentId	string		`json:"CommentId,omitempty"bson:"commentId,omitempty"`
}

type Likes struct{
	//Imageid			string			`json:"Imageid,omitempty"bson:"imageId,omitempty"`
	Username		string 			`json:"Username,omitempty"bson:"userName,omitempty"`
	TimeStamp 		time.Time 		`json:"TimeStamp,omitempty"bson:"timeStamp,omitempty"`
}

db.createUser( {user: "admin", pwd: "adminadmin", roles: [{ role: "root", db: "admin" }]});


for (var i = 101; i <= 200; i++) db.Reaction.insert({"Image_id": i,"Reaction_type":"Like","UserId":"rbd@gmail.com","Username":"ramya","Timestamp": new Date(),"Comment":"","CommentId":""})