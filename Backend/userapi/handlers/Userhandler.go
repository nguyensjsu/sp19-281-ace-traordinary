package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/dao"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/models"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/services"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/utils"
)

var user []models.User

// RegisterUserHandler creta a user
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	glog.Info("Entered the function RegisterUserEndpoint")
	var user models.Registration
	_ = json.NewDecoder(r.Body).Decode(&user)
	//Need to Remove unecessary Comments
	glog.Info("Incoming data", user)
	user.Verificationcode = utils.GenerateVerificationTocken()
	res, Message := dao.RegisterUserDao(user, r.Host)
	if res == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(Message))
	} else {
		glog.Info("Sending response", user)
		json.NewEncoder(w).Encode(user)
	}
}

//ConfirmRegistrationrHandler confirms user registration
func ConfirmRegistrationrHandler(w http.ResponseWriter, r *http.Request) {
	glog.Info("Entered the function ConfirmRegistrationrHandler")
	params := mux.Vars(r)
	var reguser models.Registration
	reguser.Userid = params["userid"]
	reguser.Verificationcode = params["verificationcode"]
	glog.Info("Incoming data", reguser)
	dao.ConfirmRegistrationDao(reguser)
}

//ResendConfirmHandler wsends confirmation mail again
func ResendConfirmHandler(w http.ResponseWriter, r *http.Request) {
	glog.Info("Entered the function ResendConfirmHandler")
	params := mux.Vars(r)
	userid := params["userid"]
	glog.Info("Incoming data", userid)
	res := dao.ResendConfirmDao(userid, r.Host)
	if !res {
		glog.Error("NOT a Valid User")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("NOT a Valid User"))
	} else {
		glog.Info("Successfully sent Resopose")
		json.NewEncoder(w).Encode(map[string]string{"result": "confirmation email resent"})
	}
}

//LoginUserHandler login specific user
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	glog.Info("Entered the function LoginUserHandler")
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	glog.Info("Incoming data", user)
	res, isValid := dao.LoginDao(user)
	if !isValid {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("NOT a Valid User"))
	} else {
		glog.Info("Sending response", res)
		json.NewEncoder(w).Encode(res)
	}
}

//ForgotPasswordrHandler will return all users
func ForgotPasswordrHandler(w http.ResponseWriter, r *http.Request) {
	glog.Info("Entered the function ForgotPasswordrHandler")
	params := mux.Vars(r)
	userid := params["userid"]
	glog.Info("Incoming data", userid)
	res := dao.ForgotpasswordDao(userid)

	if !res {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("NOT a Valid User"))
	} else {
		glog.Info("Sending response", map[string]string{"result": "success"})
		json.NewEncoder(w).Encode(map[string]string{"result": "success"})
	}
}

//DeleteUserHandler Deletes the specific user from Database
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	glog.Info("Entered the function DeleteUserHandler")
	params := mux.Vars(r)
	userid := params["userid"]
	glog.Info("Incoming data", userid)
	res := dao.DeleteUser(userid)
	if !res {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("NOT a Valid User"))
	} else {
		glog.Info("Sending response", map[string]string{"result": "success"})
		json.NewEncoder(w).Encode(map[string]string{"result": "success"})
	}
}

//GetAllUsersHandler will return all users
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	glog.Info("Entered the function GetAllUsersHandler")
	payload := dao.GetAllUsersDao()
	json.NewEncoder(w).Encode(payload)
}

//Below functions are for testing
//TestHandler for testing mailservices
func TestHandler(w http.ResponseWriter, r *http.Request) {
	glog.Info("Entered the function GetAllUsersHandler")
	var user models.Registration
	_ = json.NewDecoder(r.Body).Decode(&user)
	glog.Info("Incoming data", user)
	user.Verificationcode = utils.GenerateVerificationTocken()
	services.SendRegistrationEmail(user, r.Host)
	glog.Info("Sending response", map[string]string{"result": "This is a test mail"})
	json.NewEncoder(w).Encode(struct{ Test string }{"This is a test mail"})
}

//PingHandler for testing mailservices
func PingHandler(w http.ResponseWriter, r *http.Request) {
	glog.Info("Sending response", map[string]string{"result": "success"})
	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}
