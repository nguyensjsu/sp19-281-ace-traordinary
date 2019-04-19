package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", handlers.RegisterUserHandler).Methods("POST")
	router.HandleFunc("/users", handlers.GetAllUsersHandler).Methods("GET")
	router.HandleFunc("/users/{UserId}", handlers.ForgotPasswordrHandler).Methods("PUT")
	router.HandleFunc("/users", handlers.DeleteUserHandler).Methods("DELETE")
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
