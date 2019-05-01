package src

import (
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8004"
	}

	server := NewServer()
	server.Run(":" + port)
}
