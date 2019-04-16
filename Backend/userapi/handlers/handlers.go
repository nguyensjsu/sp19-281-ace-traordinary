package handlers

import (
	"encoding/json"
	"net/http"
	"sp19-281-ace-traordinary/Backend/userapi/src/dao"

	"github.com/gorilla/mux"
)

var user []models.User

// GetUserEndpoint gets a user
func GetUserEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	payload := dao.GetAllPeople()
	for _, p := range payload {
		if p.ID == params["id"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode("Person not found")
}

// GetAllUserEndpoint gets all user
func GetAllUserEndpoint(w http.ResponseWriter, r *http.Request) {
	payload := dao.GetAllPeople()
	json.NewEncoder(w).Encode(payload)
}

// CreateUserEndpoint creta a user
func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	dao.InsertOneValue(person)
	json.NewEncoder(w).Encode(person)
}

// DeleteUserEndpoint delets a user
func DeleteUserEndpoint(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	dao.DeletePerson(person)
}

// UpdateUserEndpoint updates a user
func UpdateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	personID := mux.Vars(r)["id"]
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	dao.UpdatePerson(person, personID)

}
