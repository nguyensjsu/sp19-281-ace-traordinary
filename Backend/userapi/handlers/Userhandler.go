package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/dao"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/models"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/services"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/utils"
)

var user []models.User

// RegisterUserHandler creta a user
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entered the function RegisterUserEndpoint")
	var user models.Registration
	_ = json.NewDecoder(r.Body).Decode(&user)
	//Need to Remove unecessary Comments
	log.Printf("Entered the function RegisterUserEndpoint")
	fmt.Println(user)
	user.Verificationcode = utils.GenerateVerificationTocken()
	res, Message := dao.RegisterUserDao(user, r.Host)
	if res == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(Message))
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

//ConfirmRegistrationrHandler confirms user registration
func ConfirmRegistrationrHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var reguser models.Registration
	reguser.Userid = params["userid"]
	reguser.Verificationcode = params["verificationcode"]
	dao.ConfirmRegistrationDao(reguser)
}

//ResendConfirmHandler wsends confirmation mail again
func ResendConfirmHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userid := params["userid"]
	res := dao.ResendConfirmDao(userid, r.Host)
	if !res {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("NOT a Valid User"))
	} else {
		json.NewEncoder(w).Encode(map[string]string{"result": "success"})
	}
}

//LoginUserHandler login specific user
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entered the function LoginUserHandler")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	log.Printf("Entered the function LoginUserHandler")
	fmt.Println(user)
	res, isValid := dao.LoginDao(user)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("NOT a Valid User"))
	} else {
		json.NewEncoder(w).Encode(res)
	}
}

//ForgotPasswordrHandler will return all users
func ForgotPasswordrHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userid := params["userid"]
	res := dao.ForgotpasswordDao(userid)
	if !res {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("NOT a Valid User"))
	} else {
		json.NewEncoder(w).Encode(map[string]string{"result": "success"})
	}
}

//DeleteUserHandler Deletes the specific user from Database
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userid := params["userid"]
	res := dao.DeleteUser(userid)
	if !res {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("NOT a Valid User"))
	} else {
		json.NewEncoder(w).Encode(map[string]string{"result": "success"})
	}
}

//GetAllUsersHandler will return all users
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	payload := dao.GetAllUsersDao()
	json.NewEncoder(w).Encode(payload)
}

//TestHandler for testing mailservices
func TestHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entered the function TestHandler")
	var user models.Registration
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.Verificationcode = utils.GenerateVerificationTocken()
	services.SendRegistrationEmail(user, r.Host)
}
