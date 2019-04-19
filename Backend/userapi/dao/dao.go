package dao

import (
	"fmt"
	"log"

	"github.com/sp19-281-ace-traordinary/Backend/userapi/models"
	"gopkg.in/mgo.v2"
)

var mongodbServer = "52.11.201.189"
var mongodbDatabase = "cmpe281"
var USERSCOLLECTION = "User"
var REGISTRATIONCOLLECTION = "Registration"

// Connect establish a connection to database
func init() {
	/**
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db = session.DB(mongodbDatabase)
	fmt.Println("Connected to MongoDB!")
	**/
}

//RegisterUserDao inserts many items from byte slice
func RegisterUserDao(user models.User) {
	fmt.Println("Entered RegisterUserDao function  ")
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(REGISTRATIONCOLLECTION)
	//hash, err := utils.EncodePassword(user.Password)
	fmt.Println(user)
	errin := c.Insert(user)
	if errin != nil {
		log.Fatal(errin)
	}
	fmt.Println("Inserted a single document: ")

}

//GetAllUsersDao returns all users in the database
func GetAllUsersDao() []models.User {
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	var results []models.User
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(USERSCOLLECTION)
	err = c.Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)
	return results
}

//ConfirmRegistrationDao once user confirms remove data from Registration and insert data to User Collection
func ConfirmRegistrationDao() {

}

//LoginDao validates weather uaer is valid or not
func LoginDao() {

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
