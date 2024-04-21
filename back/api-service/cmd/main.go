package main

import (
	"api-service/internal/router"
	"log"
)

func main() {
	r := router.SetupRouter()

	log.Println("Starting API service on port 8000...")
	if err := r.Run(":8000"); err != nil {
		log.Fatal("Failed to start API service:", err)
	}
}
