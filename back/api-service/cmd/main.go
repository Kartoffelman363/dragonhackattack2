package main

import (
	"api-service/internal/rabbitmq"
	"api-service/internal/router"
	"log"
)

func main() {
	rabbitmq.Initialize()
	defer rabbitmq.Close()

	r := router.SetupRouter()

	log.Println("Starting API service on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start API service:", err)
	}
}
