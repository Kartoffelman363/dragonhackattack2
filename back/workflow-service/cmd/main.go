package main

import (
	"context"
	"log"
	"workflow-service/internal/mongodb"
	"workflow-service/internal/router"
)

func main() {
	mongodb.Initialize()
	defer mongodb.Client.Disconnect(context.Background())

	r := router.SetupRouter()

	log.Println("Starting workflow service on port 8083...")
	if err := r.Run(":8083"); err != nil {
		log.Fatal("Failed to start API service:", err)
	}
}
