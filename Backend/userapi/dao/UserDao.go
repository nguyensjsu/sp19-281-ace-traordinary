package dao

import (
	"fmt"
	"log"

	"github.com/golang/glog"
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
	glog.Info("Entered RegisterUserDao function")
	session, err := mgo.Dial(utils.MONGODB["SERVER"])
	if err != nil {
		glog.Error("Error connecting to server")
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
		err = c.Find(bson.M{"userid": user.Userid}).One(&result)
		if err != nil {
			fmt.Println(user)
			errin := c.Insert(user)
			if errin != nil {
				glog.Error(errin)
			}
			services.SendRegistrationEmail(user, host)
			glog.Info("Successfully Regestered")
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
		glog.Error(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	var results []models.User
	c := session.DB(mongodbDatabase).C(USERSCOLLECTION)
	err = c.Find(nil).All(&results)
	if err != nil {
		glog.Error("No users found in Database", err)
	}
	glog.Error("Database Results", results)
	return results
}

//ConfirmRegistrationDao once user confirms remove data from Registration and insert data to User Collection
func ConfirmRegistrationDao(user models.Registration) (bool, models.User) {
	glog.Info("Entered ConfirmRegistrationDao function")
	var status bool
	var data models.User
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		glog.Error("Error creating session")
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(REGISTRATIONCOLLECTION)

	var result models.Registration
	query := bson.M{"userid": user.Userid}
	err = c.Find(query).One(&result)
	if err != nil {
		glog.Error("No results found")
		panic(err)
	}
	if result.Verificationcode == user.Verificationcode {
		status, data = createUserDao(result)
		if status {
			err = c.Remove(query)
			if err != nil {
				glog.Error("Remove Fail ", err)
				status = false
			}
		}
	}
	return status, data
}

func createUserDao(newuser models.Registration) (bool, models.User) {
	glog.Info("Entered CreateUserDao function")
	var user models.User
	user.Userid = newuser.Userid
	user.Password = newuser.Password
	user.Lastname = newuser.Lastname
	user.Firstname = newuser.Firstname
	user.Phonenumber = newuser.Phonenumber
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		glog.Error("Error creating session")
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
	glog.Info("Entered LoginDao function  ")
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		glog.Error("Error creating session")
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
		glog.Error("No User Found")
		panic(err)
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
	glog.Info("Entered DeleteUser function  ")
	status := true
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		glog.Error("Error creating session")
		status = false
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(mongodbDatabase).C(USERSCOLLECTION)
	query := bson.M{"userid": userid}
	err = c.Remove(query)
	if err != nil {
		glog.Error("remove fail", err)
		status = false
	}
	go services.DeleteUserEmail(userid)
	return status
}

//ForgotpasswordDao  update password of user and mail to user
func ForgotpasswordDao(userid string) bool {
	glog.Info("Entered ForgotpasswordDao function  ")
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		glog.Error("Error creating session")
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
		glog.Error("No User Found")
		return false
	}
	result.Password = utils.GenerateTemporaryPassword()
	go services.SendTemporeryPasswordEmail(result)
	err = c.Update(query, result)
	if err != nil {
		glog.Error("Error while Updating Document in Forgot Password")
		return false
	}
	return true
}

//ResendConfirmDao Resend confirmation mail and update the verification code
func ResendConfirmDao(userid string, host string) bool {
	fmt.Println("Entered ResendConfirmDao function  ")
	session, err := mgo.Dial(mongodbServer)
	if err != nil {
		glog.Error("Error creating session")
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
		glog.Info("No Registration User Found")
		return false
	}
	result.Verificationcode = utils.GenerateVerificationTocken()
	err = c.Update(query, result)
	if err != nil {
		glog.Error("Error while Updating Document in ResendConfirmDao Password")
		return false
	}
	go services.SendRegistrationEmail(result, host)
	return true
}
