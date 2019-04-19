package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sp19-281-ace-traordinary/Backend/userapi/dao"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/models"
)

var user []models.User

// RegisterUserEndpoint creta a user
func RegisterUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var user models.User
	fmt.Println("In RegisterUserEndpoint")
	_ = json.NewDecoder(r.Body).Decode(&user)
	dao.InsertOneValue(user)
	json.NewEncoder(w).Encode(user)
}

/**
// GetUserEndpoint gets a user
func GetUserEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	payload := dao.GetAllUsers()
	for _, p := range payload {
		if p.Userid == params["userid"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode("User not found")
}

// GetAllUserEndpoint gets all user
func GetAllUserEndpoint(w http.ResponseWriter, r *http.Request) {
	payload := dao.GetAllUsers()
	json.NewEncoder(w).Encode(payload)
}

// RegisterUserEndpoint creta a user
func RegisterUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var user models.User
	fmt.Println("In RegisterUserEndpoint")
	_ = json.NewDecoder(r.Body).Decode(&user)
	dao.InsertOneValue(user)
	json.NewEncoder(w).Encode(user)
}

// DeleteUserEndpoint delets a user
func DeleteUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	dao.DeleteUser(user)
}

/**func UpdateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	personID := mux.Vars(r)["id"]
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&person)
	dao.UpdatePerson(person, personID)

}**/
