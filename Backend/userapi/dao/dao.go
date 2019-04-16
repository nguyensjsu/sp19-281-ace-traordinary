package dao

import (
	"context"
	"fmt"
	"log"

	"github.com/sp19-281-ace-traordinary/Backend/userapi/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CONNECTIONSTRING DB connection string
const CONNECTIONSTRING = "mongodb://cmpe281:cmpe281@ds139896.mlab.com:39896/cmpe281"

// DBNAME Database name
const DBNAME = "cmpe281"

// COLLNAME Collection name
const COLLNAME = "user"

var db *mongo.Database

// Connect establish a connection to database
func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI(CONNECTIONSTRING))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	// Collection types can be used to access the database
	db = client.Database(DBNAME)
}

// InsertManyValues inserts many items from byte slice
func InsertManyValues(user []models.User) {
	var ppl []interface{}
	for _, p := range user {
		ppl = append(ppl, p)
	}
	_, err := db.Collection(COLLNAME).InsertMany(context.Background(), ppl)
	if err != nil {
		log.Fatal(err)
	}
}

// InsertOneValue inserts one item from Person model
func InsertOneValue(user models.User) {
	fmt.Println(user)
	_, err := db.Collection(COLLNAME).InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllUsers returns all users from DB
func GetAllUsers() []models.User {
	cur, err := db.Collection(COLLNAME).Find(context.Background(), nil, nil)
	if err != nil {
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

// UpdatePerson updates an existing person
