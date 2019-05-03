package main

import (
	"os"
	//"fmt"
	//"database/sql"
	// "github.com/go-sql-driver/mysql"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8005"
	}

	server := NewServer()
	server.Run(":" + port)
}
