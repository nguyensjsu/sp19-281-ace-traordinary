/*
	picsqueryapi REST API (Version 1)
*/

package main

import (
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8002"
	}

	server := NewServer()
	server.Run(":" + port)
}
