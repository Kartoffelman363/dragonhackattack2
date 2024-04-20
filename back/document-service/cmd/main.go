package main

import (
	"context"
	"document-service/internal/mongodb"
)

func main() {
	mongodb.Initialize()
	defer mongodb.Client.Disconnect(context.Background())

}
