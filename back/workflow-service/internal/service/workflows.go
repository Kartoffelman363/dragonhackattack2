package service

import (
	"context"
	"fmt"
	"time"
	"workflow-service/internal/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllWorflows() ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := mongodb.GetCollection("yourDatabase", "worflows")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}
	bsonBytes, err := bson.Marshal(results)
	if err != nil {
		return nil, err
	}
	return bsonBytes, nil
}

func GetWorkflowByID(id string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	collection := mongodb.GetCollection("yourDatabase", "worflows")
	var result bson.M
	if err := collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&result); err != nil {
		return nil, err
	}
	bsonBytes, err := bson.Marshal(result)
	if err != nil {
		return nil, err
	}
	return bsonBytes, nil
}

func DeleteWorkflow(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid ID format: %v", err)
	}

	collection := mongodb.GetCollection("yourDatabase", "worflows")
	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}
