package main

import (
	"context"
	"workflow-service/internal/mongodb"
)

func main() {
	mongodb.Initialize()
	defer mongodb.Client.Disconnect(context.Background())

}
