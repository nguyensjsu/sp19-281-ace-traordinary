package main

import (
	"os"
	"net/http",
	"log"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
  InitialMigration()
	server := NewServer()
	server.Run(":" + port)
}
