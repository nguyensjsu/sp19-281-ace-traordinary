//package main

/**
import (
	"fmt"
	"net/http"
	"os"
)

func main(){

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8000"
	}
	Pool := newPool()
	go Pool.run()
	//http.HandleFunc("/", serveHome)
	fmt.Println("http server started on :8000")
	http.HandleFunc("/ws/{Id}", func(w http.ResponseWriter, r *http.Request) {
		serveWs(Pool, w, r)
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
	panic( err)
	}

	//http.HandleFunc("/ws", handleConnections)
	//go handleMessages()
	//err := http.ListenAndServe(":8000", nil)
	//if err != nil {
	//	fmt.Errorf("ListenAndServe: ", err)
//	}

}
**/
