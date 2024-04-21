package main

import (
	"context"
	"document-service/internal/mongodb"
	"document-service/internal/router"
	"log"
)

func main() {
	mongodb.Initialize()
	defer mongodb.Client.Disconnect(context.Background())

	r := router.SetupRouter()

	log.Println("Starting document service on port 8002...")
	if err := r.Run(":8002"); err != nil {
		log.Fatal("Failed to start API service:", err)
	}
}
