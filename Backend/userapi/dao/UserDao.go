package dao

import (
	"fmt"
	"log"

	"github.com/sp19-281-ace-traordinary/Backend/userapi/models"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
func RegisterUserDao(user models.Registration) (bool, string) {
	fmt.Println("Entered RegisterUserDao function  ")
	session, err := mgo.Dial(utils.MONGODB["SERVER"])
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(utils.MONGODB["DATABASE"]).C(utils.MONGODB["SERVER"])
	var result bson.M
	//Checking if the new user is already present in user table
	err = c.Find(bson.M{"userid": user.Userid}).One(&result)
	if err != nil {
		c = session.DB(mongodbDatabase).C(utils.MONGODB["REGISTRATIONCOLLECTION"])
		//hash, err := utils.EncodePassword(user.Password)
		fmt.Println(user.Userid)
		err = c.Find(bson.M{"userid": user.Userid}).One(&result)
		if err != nil {
			fmt.Println(user)
			errin := c.Insert(user)
			if errin != nil {
				log.Fatal(errin)
			}
			fmt.Println("Successfully Regestered")
		} else {
			return false, "Already In Registration Table"
		}
	} else {
		return false, "User is Already Present"
	}
	return true, "Successfully Regestered"
}

//GetAllUsersDao returns all users in the database
func GetAllUsersDao() []models.User {
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var results []models.User
	c := session.DB(mongodbDatabase).C(USERSCOLLECTION)
	err = c.Find(nil).All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)
	return results
}

//ConfirmRegistrationDao once user confirms remove data from Registration and insert data to User Collection
func ConfirmRegistrationDao(user models.Registration) (bool, models.User) {
	fmt.Println("Entered LoginDao function  ")
	var status bool
	var data models.User
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(REGISTRATIONCOLLECTION)

	var result models.Registration
	query := bson.M{"userid": user.Userid}
	err = c.Find(query).One(&result)
	if err != nil {
		panic(err)
	}
	if result.Verificationcode == user.Verificationcode {
		status, data = createUserDao(result)
		if status {
			err = c.Remove(query)
			if err != nil {
				fmt.Printf("remove fail %v\n", err)
				status = false
			}
		}
	}
	return status, data
}

func createUserDao(newuser models.Registration) (bool, models.User) {
	fmt.Println("Entered LoginDao function  ")
	var user models.User
	user.Userid = newuser.Userid
	user.Password = newuser.Password
	user.Lastname = newuser.Lastname
	user.Firstname = newuser.Firstname
	user.Phonenumber = newuser.Phonenumber
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(USERSCOLLECTION)
	errin := c.Insert(user)
	if errin != nil {
		panic(err)
	}
	return true, user
}

//LoginDao validates weather uaer is valid or not
func LoginDao(user models.User) models.User {
	fmt.Println("Entered LoginDao function  ")
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(USERSCOLLECTION)

	var result models.User
	//Checking if the new user is already present in user table
	err = c.Find(bson.M{"userid": user.Userid}).One(&result)
	if err != nil {
		log.Println("No User Found")
	}
	return result
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