package main

import (
	"fmt"
	"log"
	"net/http"

	"./userapi/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.GetAllPeopleEndpoint).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/users", handlers.CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/users", handlers.DeletePersonEndpoint).Methods("DELETE")
	router.HandleFunc("/users/{id}", handlers.UpdateUserEndpoint).Methods("PUT")
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
