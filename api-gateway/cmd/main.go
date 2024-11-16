package main

import (
	"api-gateway/cmd/bootstrap"
	"log"
)

func main() {
	router := bootstrap.NewRouter()

	if err := router.Run(":5000"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
