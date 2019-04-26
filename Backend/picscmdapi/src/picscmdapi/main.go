/*
	picscmdapi REST API (Version 1)
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	headersOk := h.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := h.AllowedOrigins([]string{"*"})
	methodsOk := h.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	router.HandleFunc("/images", UploadPictureHandler).Methods("POST")

	fmt.Println("Starting server on port 3001...")
	log.Fatal(http.ListenAndServe(":3001", h.CORS(headersOk, methodsOk, originsOk)(router)))
}
