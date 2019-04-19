package dao

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"

	"github.com/sp19-281-ace-traordinary/Backend/userapi/models"
	"gopkg.in/mgo.v2"
)

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://cmpe281:cmpe281@ds139896.mlab.com:39896/cmpe281"

var mongodbServer = "mongodb+srv://cmpe281:cmpe281@cluster0-p8lxi.mongodb.net/test?retryWrites=true"
var mongodbDatabase = "cmpe281"
var mongodbCollection = "user"

// DBNAME Database name
const DBNAME = "cmpe281"

// COLLNAME Collection name
const COLLNAME = "User"

var db *mgo.Database

// Connect establish a connection to database
func init() {
	dialInfo, err := mgo.ParseURL(mongodbServer)
	tlsConfig := &tls.Config{}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db = session.DB(mongodbDatabase)
	fmt.Println("Connected to MongoDB!")

}

// InsertOneValue inserts many items from byte slice
func InsertOneValue(user models.User) {
	fmt.Println("In InsertOneValue")
	fmt.Println(db)
	collection := db.C(COLLNAME)
	fmt.Println(collection)
	fmt.Println("Successfully go collection")
	err := collection.Insert(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ")

}

/**
// InsertManyValues inserts many items from byte slice
func InsertManyValues(user []models.User) {
	var ppl []interface{}
	for _, p := range user {
		ppl = append(ppl, p)
	}
	_, err := db.C(COLLNAME).InsertMany(context.Background(), ppl)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertOneValue inserts one item from Person model
func InsertOneValue(user models.User) {
	fmt.Println("In InsertOneValue")
	fmt.Println(db)
	collection := db.Collection(COLLNAME)
	fmt.Println(collection)
	fmt.Println("Successfully go collection")
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

}

// GetAllUsers returns all users from DB
func GetAllUsers() []models.User {
	cur, err := db.Collection(COLLNAME).Find(context.Background(), nil, nil)
	if err != nil {
		log.Fatal("Exception in GetAllUsers")
		log.Fatal(err)
	}
	var elements []models.User
	var elem models.User
	// Get the next result from the cursor
	for cur.Next(context.Background()) {
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		elements = append(elements, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.Background())
	return elements
}

// DeletePerson deletes an existing user
func DeleteUser(user models.User) {
	_, err := db.Collection(COLLNAME).DeleteOne(context.Background(), user, nil)
	if err != nil {
		log.Fatal(err)
	}
}
**/
// UpdatePerson updates an existing person
