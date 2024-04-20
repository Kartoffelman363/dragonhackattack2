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

	log.Println("Starting document service on port 8082...")
	if err := r.Run(":8082"); err != nil {
		log.Fatal("Failed to start API service:", err)
	}
}
