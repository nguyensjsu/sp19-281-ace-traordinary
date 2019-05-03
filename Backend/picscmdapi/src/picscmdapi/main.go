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
	methodsOk := h.AllowedMethods([]string{"HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	router.HandleFunc("/images", UploadPictureHandler).Methods("POST")
	router.HandleFunc("/images/:imageid", UpdatePictureHandler).Methods("PUT")
	router.HandleFunc("/images/{imageid}", DeletePictureHandler).Methods("DELETE")
	router.HandleFunc("/picscmd/ping", PingHandler).Methods("GET")
	fmt.Println("Starting server on port 8001...")
	log.Fatal(http.ListenAndServe(":8001", h.CORS(headersOk, methodsOk, originsOk)(router)))
}
