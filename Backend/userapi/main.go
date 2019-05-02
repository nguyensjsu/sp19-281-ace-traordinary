package main

import (
	"fmt"
	"log"
	"net/http"

	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sp19-281-ace-traordinary/Backend/userapi/handlers"
)

func main() {
	router := mux.NewRouter()
	headersOk := h.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := h.AllowedOrigins([]string{"*"})
	methodsOk := h.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router.HandleFunc("/users", handlers.RegisterUserHandler).Methods("POST")
	router.HandleFunc("/user", handlers.LoginUserHandler).Methods("POST")
	router.HandleFunc("/user", handlers.ConfirmRegistrationrHandler).Methods("GET")
	router.HandleFunc("/users", handlers.ResendConfirmHandler).Methods("POST")
	router.HandleFunc("/users/{userid}", handlers.ForgotPasswordrHandler).Methods("GET")
	router.HandleFunc("/users", handlers.DeleteUserHandler).Methods("DELETE")
	//Below routers are not into the business they are for testing
	router.HandleFunc("/test", handlers.TestHandler).Methods("POST")
	router.HandleFunc("/ping", handlers.PingHandler).Methods("GET")

	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", h.CORS(headersOk, methodsOk, originsOk)(router)))
}
