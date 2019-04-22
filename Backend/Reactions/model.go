package src

import "time"

type Reaction struct{
	Id				string 		`bson:"_id, omitempty"`
	Image_id 		string   	`bson:"Image_id, omitempty"`
	Reaction_type 	string		`json:"ReactionType,omitempty" bson:"Reaction_type, omitempty"`
	UserId			string 		`json:"UserId,omitempty"`
	Username 		string 	 	`json:"Username,omitempty"`
	Timestamp 		time.Time	`json:"Timestamp,omitempty"`
	Comment 		string  	`json:"Comment,omitempty"`
}

type ImageReaction struct {
	Imageid 	string
	likes	 	[]Likes
	comments	[]Comments

}

type Comments struct{
	Username	string 		`json:"Username,omitempty"	bson:"Username,omitempty"`
	Comment		string 		`json:"Comment,omitempty"	bson:"Comment,omitempty"`
	TimeStamp 	time.Time 	`json:"TimeStamp,omitempty"	bson:"TimeStamp,omitempty"`
}

type Likes struct{
	Username		string 			`json:"Username,omitempty"	bson:"Username,omitempty"`
	TimeStamp 		time.Time 		`json:"TimeStamp,omitempty"	bson:"TimeStamp,omitempty"`
}