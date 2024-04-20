package main

import (
	openai_client "coderunner-service/internal/openai"
	"coderunner-service/internal/router"
	"log"
)

func main() {
	openai_client.Initialize()

	r := router.SetupRouter()

	log.Println("Starting coderunner service on port 8081...")
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Failed to start API service:", err)
	}
}
