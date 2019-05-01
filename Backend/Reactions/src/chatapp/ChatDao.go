package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

var mongodbServer = "54.241.227.172"
var mongodbDatabase = "cmpe281"
var Messages_Collection = "Conversations"
var Unread_Messages = "UnreadCollection"


func storeMessage(message Message) (error,bool) {
	fmt.Println("Entered LoginDao function  ")

	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(Messages_Collection)

	err = c.Insert(message)
	if err!=nil {
		log.Println( "Error: ", err )

		return err,false
	}

	return nil,true
}
func storeUnreadMessage(message Message) (error,bool) {
	fmt.Println("Entered LoginDao function  ")

	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(Unread_Messages)

	err = c.Insert(message)
	if err!=nil {
		log.Println( "Error: ", err )

		return err,false
	}

	return nil,true
}

func readUnreadMessages(clientid string) ([]Message,bool) {
	fmt.Println("Entered LoginDao function  ")

	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(Unread_Messages)

	fmt.Println( "Fetching unread messages for: ", clientid )

	var messages []Message
	err = c.Find(bson.M{"Receiverid": clientid}).All(&messages)
	if err!=nil {
		fmt.Println( "Error: ", err )
		return nil,false
	}
	if len(messages)==0 {
		fmt.Println( "No such document: " )
		return nil,false
	}
	return messages,true
}

func removeUnreadMessages(clientid string) (error,bool) {
	fmt.Println("Entered LoginDao function  ")

	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(Unread_Messages)

	fmt.Println( "Deleting unread messages for: ", clientid )

	query := bson.M{"Receiverid":clientid}

	errin := c.Remove(query)
	if errin != nil {
		panic(err)
		return err,false
	}

	return nil,true
}

func updateConversations(messageId string, status string, updated time.Time) (error, bool){


	fmt.Println("Entered LoginDao function  ")

	session, err := mgo.Dial(mongodbServer)
	if err != nil {
	panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(Messages_Collection)

		fmt.Println( "Updating status messages for: ", messageId )

	query := bson.M{
	"MessageId": messageId,

		}

	update := bson.M{
	"$set": bson.M{
	"Status": status,
	"Lastupdated": updated,
		}}

		err= c.Update(query,update)


		if err != nil {
			panic(err)
			return err,false
		}

	return nil,true

}

func loadConverstaion(userid string, receiverid string) ([]Message, bool){

//var limit = 5;
	fmt.Println("Entered LoginDao function  ")

	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(Messages_Collection)
	var messages []Message
	fmt.Println( "Updating status messages for: ", userid )
	users := []string{userid,receiverid}
	//matchQuery := bson.M{"UserId":userid,"Receiverid":"receiverid"}
	match := bson.M{"$and" : []bson.M{bson.M{ "UserId": bson.M{"$in" : users}},bson.M{"Receiverid" : bson.M{"$in" : users}}}}
	sortQuery1 := bson.M{"Time": -1 }
	sortQuery2 := bson.M{"Time": 1 }



	pipeline := []bson.M{
		{"$match": match},
		{"$sort" : sortQuery1},
		{"$limit" : 5},
		{"$sort": sortQuery2},
	}

	err = c.Pipe(pipeline).All(&messages)
	if err!=nil {
		fmt.Println( "Error: ", err )
		return nil,false
	}

	return messages,true

}