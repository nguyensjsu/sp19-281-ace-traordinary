package dao

import (
	"fmt"
	"log"

	"github.com/sp19-281-ace-traordinary/Backend/userapi/models"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/services"
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
func RegisterUserDao(user models.Registration, host string) (bool, string) {
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
			services.SendRegistrationEmail(user, host)
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
	fmt.Println("Entered ConfirmRegistrationDao function  ")
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
func LoginDao(user models.User) (models.User, bool) {
	fmt.Println("Entered LoginDao function  ")
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(USERSCOLLECTION)
	var result models.User
	query := bson.M{"userid": user.Userid}
	//Checking if the new user is already present in user table
	err = c.Find(query).One(&result)
	if err != nil {
		log.Println("No User Found")
	}
	res := validatePassword(user.Password, result.Password)
	if !res {
		return result, false
	}
	return result, true
}
func validatePassword(in string, dbpassword string) bool {
	if in == dbpassword {
		return true
	}
	return false
}

//DeleteUser removes user with userID from the database
func DeleteUser(userid string) bool {
	status := true
	fmt.Println("Entered DeleteUser function  ")
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		log.Panic(err)
		status = false
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(USERSCOLLECTION)
	query := bson.M{"userid": userid}
	err = c.Remove(query)
	if err != nil {
		fmt.Printf("remove fail %v\n", err)
		status = false
	}
	go services.DeleteUserEmail(userid)
	return status
}

//ForgotpasswordDao  update password of user and mail to user
func ForgotpasswordDao(userid string) bool {
	fmt.Println("Entered LoginDao function  ")
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		log.Panic(err)
		return false
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(USERSCOLLECTION)
	var result models.User
	query := bson.M{"userid": userid}
	err = c.Find(query).One(&result)
	if err != nil {
		log.Println("No User Found")
		return false
	}
	result.Password = utils.GenerateTemporaryPassword()
	go services.SendTemporeryPasswordEmail(result)
	err = c.Update(query, result)
	if err != nil {
		log.Println("Error while Updating Document in Forgot Password")
		return false
	}
	return true
}

//ResendConfirmDao Resend confirmation mail and update the verification code
func ResendConfirmDao(userid string, host string) bool {
	fmt.Println("Entered ResendConfirmDao function  ")
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		log.Panic(err)
		return false
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(REGISTRATIONCOLLECTION)
	var result models.Registration
	query := bson.M{"userid": userid}
	err = c.Find(query).One(&result)
	if err != nil {
		log.Println("No Registration User Found")
		return false
	}
	result.Verificationcode = utils.GenerateVerificationTocken()
	err = c.Update(query, result)
	if err != nil {
		log.Println("Error while Updating Document in ResendConfirmDao Password")
		return false
	}
	go services.SendRegistrationEmail(result, host)
	return true
}
