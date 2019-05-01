package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"sort"
	"strings"
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
	convId := getConversationId(userid,receiverid)
	var messages []Message
	fmt.Println( "Updating status messages for: ", userid )
	//users := []string{userid,receiverid}
	//matchQuery := bson.M{"UserId":userid,"Receiverid":"receiverid"}
	match := bson.M{"ConversationId":convId}
	sortQuery1 := bson.M{"Time": -1 }
	//sortQuery2 := bson.M{"Time": 1 }



	pipeline := []bson.M{
		{"$match": match},
		{"$sort" : sortQuery1},
		{"$limit" : 5},
	}

	err = c.Pipe(pipeline).All(&messages)
	if err!=nil {
		fmt.Println( "Error: ", err )
		return nil,false
	}

	return messages,true

}

func updateSeenStatus(sender string,recver string, time2 time.Time) (string,bool){



	fmt.Println("Entered LoginDao function  ")

	session, err := mgo.Dial(mongodbServer)
	if err != nil {
	panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(Messages_Collection)

	fmt.Println( "Updating seen by the user : ",recver  )

	query := bson.M{
	"UserId": sender,
	"Receiverid" : recver,
	}

	update := bson.M{
	"$set": bson.M{
	"Status": "Seen",
	"Lastupdated": time2,
	}}

			err= c.Update(query,update)


if err != nil {
panic(err)
return "",false
}

return getConversationId(sender,recver),true


}

func getConversationId(user1 string, user2 string) string{

	arrayId := []string{strings.ToLower(user1),strings.ToLower(user2)}
	sort.Strings(arrayId);
	uniqueId := strings.Join(arrayId,"")
	fmt.Println("conv id generated is ",uniqueId)
	return uniqueId


}